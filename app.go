package main

import (
	"fmt"
	"time"

	"github.com/wangkebin/kbds-client/service"
)

func main() {

	cfg, err := service.LoadConfig(".")
	if err != nil {
		fmt.Printf("%v", err)
	}

	db, err := service.Connect(cfg.DbConnStr, cfg.Debug)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	start := time.Now()
	service.Walk(&cfg, db)
	fmt.Println("time taken (ms): ", time.Since(start).Milliseconds())

}
