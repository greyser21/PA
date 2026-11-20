package main

import (
	"log"

	"upcycleconnect/internal/handlers"
	"upcycleconnect/internal/store"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	s := store.New()
	api := handlers.New(s)
	api.Register(r)

	// Serve frontend static files (HTML, CSS, JS)
	r.Static("/css", "./web/css")
	r.Static("/js", "./web/js")
	r.Static("/fonts", "./web/fonts")
	r.StaticFile("/logo.jpg", "./web/logo.jpg")
	r.StaticFile("/", "./web/index.html")
	r.StaticFile("/index.html", "./web/index.html")
	r.StaticFile("/user.html", "./web/user.html")
	r.StaticFile("/pro.html", "./web/pro.html")
	r.StaticFile("/employee.html", "./web/employee.html")
	r.StaticFile("/admin.html", "./web/admin.html")

	log.Println("UpcycleConnect server listening on http://localhost:8080")
	if err := r.Run(":8090"); err != nil {
		log.Fatal(err)
	}
}
