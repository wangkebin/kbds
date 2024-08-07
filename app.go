package main

import (
	"fmt"
	"kbds/models"

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

	run(cfg.StartPath, db)


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
