package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/entrik/kathisto/renderer"
	"github.com/entrik/kathisto/service"
)

var (
	version   = "0.5"
	userAgent = fmt.Sprintf("kathisto/%s", version)
	pubDir    = getEnv("PUBLIC_ADDR", "/dist")
	host      = getEnv("HOST", "")
	port      = getEnv("PORT", "80")
)

func main() {
	log.Printf("kathisto v%s - Server-Side rendering with Go/Headless Chrome\n", version)

	// Create a Chrome Renderer and bind the Render func to /
	r := renderer.NewChromeRenderer(userAgent)
	rs := service.NewService(r, host, pubDir)
	http.HandleFunc("/", rs.Render)

	// Always run a basic http server
	log.Println("Listening on port :", port)
	http.ListenAndServe(":"+port, nil)
}

// getEnv retrieves variables from the environment and falls
// back to a passed fallback variable if it isn't set
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
