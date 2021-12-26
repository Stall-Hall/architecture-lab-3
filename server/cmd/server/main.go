package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"architecture-lab-3/server/db"
)

var httpPortNumber = flag.Int("p", 8080, "HTTP port number")

func NewDbConnection() (*sql.DB, error) {
	conn := &db.Connection{
		DbName:     "lab3",
		User:       "postgres",
		Password:   "qwe",
		Host:       "localhost",
		DisableSSL: true,
	}
	return conn.Open()
}

func main() {
	flag.Parse()

	if server, err := ComposeApiServer(HttpPortNumber(*httpPortNumber)); err == nil {
		go func() {
			log.Println("Starting server...")

			err := server.Start()
			if err == http.ErrServerClosed {
				log.Printf("HTTP server stopped")
			} else {
				log.Fatalf("Cannot start HTTP server: %s", err)
			}
		}()

		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt)
		<-sigChannel

		if err := server.Stop(); err != nil && err != http.ErrServerClosed {
			log.Printf("Error stopping the server: %s", err)
		}
	} else {
		log.Fatalf("Cannot initialize server: %s", err)
	}
}
