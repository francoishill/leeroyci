package main

import (
	"log"
	"net/http"
	"os"

	"github.com/francoishill/leeroyci/database"
	"github.com/francoishill/leeroyci/runner"
	"github.com/francoishill/leeroyci/web"
	"github.com/francoishill/leeroyci/websocket"
)

func main() {
	database.NewDatabase("", "")
	websocket.NewServer()
	go runner.Runner()

	router := web.Routes()
	config := database.GetConfig()

	if config.Cert != "" {
		log.Fatalln(http.ListenAndServeTLS(port(), config.Cert, config.Key, router))
	} else {
		log.Fatalln(http.ListenAndServe(port(), router))
	}
}

func port() string {
	port := os.Getenv("PORT")

	if port == "" {
		return ":8082"
	}

	return ":" + port
}
