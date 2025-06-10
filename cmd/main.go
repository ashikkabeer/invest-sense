package main

import (
	"log"

	"github.com/ashikkabeer/invest-sense/internal/server"
)

func main() {
	log.Println("Starting server...")
	e := server.NewServer()

	// repo := repository.NewUserRepository()
	// svc := service.NewUserService(repo)
	// h := handler.NewHandler(svc)

	// server.RegisterRoutes(e, h)

	e.Logger.Fatal(e.Start(":8080"))
}
