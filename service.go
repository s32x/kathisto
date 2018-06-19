package main

import (
	"log"
	"net/http"
	"strings"
)

// Servicer defines all functionality for prerender web-pages
type Servicer interface {
	Prerender(w http.ResponseWriter, r *http.Request)
}

// service contains all dependencies needed for a Servicer
type service struct {
	renderer   Renderer
	strictHost string
	pubDir     string
}

// NewService generates and returns a new Servicer
func NewService(renderer Renderer, strictHost, pubDir string) Servicer {
	return &service{renderer, strictHost, pubDir}
}

// Prerender handles all requests to the static frontend. Depending on the
// requested resource and the clients user-agent, we determine whether to serve
// the static resource or a prerendered full html page
func (s *service) Prerender(w http.ResponseWriter, r *http.Request) {
	if s.strictHost != "" && !strings.Contains(r.Host, s.strictHost) {
		return
	}
	log.Printf("New valid request on Host - %s from IP - %s\n", r.Host, r.RemoteAddr)

	r.URL.Scheme = "http"
	if r.TLS != nil {
		r.URL.Scheme = "https"
	}
	r.URL.Host = r.Host

	path := r.URL.Path // The path requested from our domain
	log.Printf("Handling request for URL - %s\n", r.URL.String())

	// Adding powered by header
	w.Header().Add("X-Powered-By", "Kathisto")

	// If a static file
	if strings.Contains(path, ".") {
		log.Printf("Serving static file %s - %s\n", s.pubDir+path, r.UserAgent())
		http.ServeFile(w, r, s.pubDir+path)
		return
	}

	// If a NoOp or something with a querystring
	if s.renderer.IsNoOp(r.URL.Query()) || strings.Contains(r.URL.String(), "?") {
		log.Printf("Serving static index %s - %s\n", s.pubDir+"/index.html", r.UserAgent())
		http.ServeFile(w, r, s.pubDir+"/index.html")
		return
	}

	// Attempt to render
	log.Printf("Attempting to render %s - %s\n", s.pubDir+path, r.UserAgent())
	if body, err := s.renderer.Render(r.URL.String()); err == nil {
		w.Write(body)
	} else {
		log.Printf("Falling back to static index %s - %s\n", s.pubDir+"/index.html", r.UserAgent())
		http.ServeFile(w, r, s.pubDir+"/index.html")
	}
}
