package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	version         = "0.2"
	cacheExpiration = 12 * time.Hour
)

func main() {
	log.Printf("Kathisto v%s - Server-Side rendering with Go/PhantomJS\n", version)
	pubDir := os.Getenv("PUBLIC_DIR")
	if pubDir == "" {
		pubDir = "/dist"
	}

	// Create a PhantomJS renderer and attach the prerender func to /
	r := NewPJSRenderer(cacheExpiration, fmt.Sprintf("Kathisto/%s", version))
	rs := NewService(r, os.Getenv("STRICT_HOST"), pubDir)
	http.HandleFunc("/", rs.Prerender)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	// Always run a basic http server
	log.Println("Listening on port :", port)
	http.ListenAndServe(":"+port, nil)
}
