package main

import (
	"fmt"
	"io/fs"
	"kbds/models"
	"path/filepath"

	"github.com/spf13/viper"
)

func main() {

	cfg, err := LoadConfig(".")
	if err != nil {
		fmt.Printf("%v", err)
	}

	db, err := Connect(cfg.DbConnStr)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	filepath.WalkDir(cfg.StartPath, func(path string, d fs.DirEntry, err error) error {
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

		if !d.IsDir() {
			db.Create(&f)
		}
		return nil
	})

}

func LoadConfig(path string) (config models.Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%v", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("%v", err)
	}
	return
}
