package main

import (
	"log"

	"upcycleconnect/api/internal/config"
	"upcycleconnect/api/internal/router"
	"upcycleconnect/api/internal/store"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	gin.SetMode(cfg.GinMode)
	r := gin.Default()
	st := store.New()
	router.Register(r, st, cfg)

	log.Printf("UpcycleConnect API listening on %s", cfg.Addr)
	if err := r.Run(cfg.Addr); err != nil {
		log.Fatal(err)
	}
}
