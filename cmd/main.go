package main

import (
	"log"

	"github.com/alexvlasov182/todo"
	"github.com/alexvlasov182/todo/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	println("Hello World!")
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
