package main

import (
	"log"
	"url-shortener/internal/handler"

	"github.com/pkg/errors"
)

func main() {
	handler, err := handler.NewHandler()
	if err != nil {
		log.Fatal(errors.Wrap(err, "handler.NewHandler"))
	}
	handler.Router.Run(":8000")
}
