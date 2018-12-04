package api /* import "s32x.com/kathisto/api" */

import (
	"log"
	"net/http"

	"s32x.com/kathisto/renderer"
	"s32x.com/kathisto/service"
)

// Start starts the kathisto API service using the passed params
func Start(version, userAgent, pubDir, host, port string) {
	log.Printf("kathisto v%s - Server-Side rendering with Go/Headless Chrome\n", version)

	// Create a Chrome Renderer and bind the Render func to /
	r := renderer.NewChromeRenderer(userAgent)
	rs := service.NewService(r, host, pubDir)
	http.HandleFunc("/", rs.Render)

	// Always run a basic http server
	log.Println("Listening on port :", port)
	http.ListenAndServe(":"+port, nil)
}
