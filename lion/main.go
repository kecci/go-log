package main

import (
	"net/http"

	"github.com/celrenheit/lion/middleware"
)

func main() {
	log := middleware.NewLogger()
	log.ServeNext(http.NotFoundHandler())
	// Not logger as logger (lion is middleware...)
}