package main

import (
	"io/fs"
	"kbds/models"
	"path/filepath"

	//"runtime"
	"sync"

	"gorm.io/gorm"
)

func collector(fmetas <-chan models.FMeta, db *gorm.DB) error {
	filemetas := make([]models.FMeta, 0)
	for f := range fmetas {
		filemetas = append(filemetas, f)
		if len(filemetas) > 999 {
			res := db.CreateInBatches(&filemetas, 1000)
			if res.Error != nil {
				return res.Error
			}
			filemetas = make([]models.FMeta, 0)
		}
	}
	res := db.CreateInBatches(&filemetas, 1000)
	if res.Error != nil {
		return res.Error
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

func run(startPath string, db *gorm.DB) models.Results {
	//workers := 2 * runtime.GOMAXPROCS(0)
	fmetas := make(chan models.FMeta)
	//done := make(chan bool)
	var wg sync.WaitGroup

	go traversal(startPath, fmetas)
	close(fmetas)
	wg.Add(1)
	go func() {
		collector(fmetas, db)
		wg.Done()
	}()
	wg.Wait()
	// for i := 0; i < workers; i++ {
	// 	go collectHashes(fmetas, done)
	// }
	return nil
}
