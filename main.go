package main

import (
	"fmt"
	"github.com/umerm-work/xm-exercise/config"
	"github.com/umerm-work/xm-exercise/infrastructure/db"
	"github.com/umerm-work/xm-exercise/infrastructure/router"
	"log"
)

func main() {
	conf := config.Load()
	db.Connect(conf)
	r := router.Setup()
	addr := fmt.Sprintf(":%s", conf.AppPort)

	if err := r.Run(addr); err != nil {
		log.Fatalln("Failed to start the application.", err)
	}
}
