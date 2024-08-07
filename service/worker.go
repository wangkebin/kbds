package service

import (
	"io/fs"
	"path/filepath"

	"github.com/go-faster/errors"
	"github.com/wangkebin/kbds-client/models"

	//"runtime"
	"sync"

	"gorm.io/gorm"
)

func collector(fmetas <-chan models.FMeta, batchsize int, db *gorm.DB) error {
	filemetas := make([]models.FMeta, 0)
	for f := range fmetas {
		filemetas = append(filemetas, f)
		if len(filemetas) >= batchsize {
			if err := CreateInBatches(db, &filemetas, batchsize); err != nil {
				return errors.Wrap(err, "failed to create fmeta records in db")
			}
			filemetas = make([]models.FMeta, 0)
		}
	}
	if err := CreateInBatches(db, &filemetas, batchsize); err != nil {
		return errors.Wrap(err, "failed to create fmeta records in db")
	}
	return nil
}

func traversal(startPath string, fmetas chan<- models.FMeta) error {
	visit := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		f := models.FMeta{
			Loc:  path,
			Size: info.Size(),
			Name: d.Name(),
			Ext:  "",
		}
		fmetas <- f

		return nil
	}
	return filepath.WalkDir(startPath, visit)
}

func Walk(cfg *models.Config, db *gorm.DB) models.Results {
	//workers := 2 * runtime.GOMAXPROCS(0)
	fmetas := make(chan models.FMeta)
	//done := make(chan bool)
	var wg sync.WaitGroup
	
	wg.Add(1)
	go func() {
		traversal(cfg.StartPath, fmetas)
		close(fmetas)
	}()

	go func() {
		collector(fmetas, cfg.BatchSize, db)
		wg.Done()
	}()
	wg.Wait()
	// for i := 0; i < workers; i++ {
	// 	go collectHashes(fmetas, done)
	// }
	return nil
}
