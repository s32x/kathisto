package service

import (
	"net/http"
	"strings"

	"github.com/starboy/kathisto/renderer"
)

// Service contains all dependencies needed for a Servicer
type Service struct {
	renderer     renderer.Renderer
	host, pubDir string
}

// NewService generates and returns a new Servicer
func NewService(renderer renderer.Renderer, host, pubdir string) *Service {
	return &Service{renderer, host, pubdir}
}

// Render handles all requests to our webapp. Depending on
// the requested resource we determine whether to serve the
// static resource or a fully rendered html page
func (s *Service) Render(w http.ResponseWriter, r *http.Request) {
	// Set the scheme and hostname on the URL if relative
	if !r.URL.IsAbs() {
		r.URL.Scheme = "http"
		if r.TLS != nil {
			r.URL.Scheme = "https"
		}
		r.URL.Host = r.Host
	}

	// Add the X-Powered-By header
	w.Header().Add("X-Powered-By", "kathisto")

	// Don't return anything if the hosts don't match
	if insContains(r.URL.Host, s.host) {
		w.Write([]byte("Invalid host"))
		return
	}

	// If a static file is requested (ex: sitemap.xml or robots.txt)
	// serve the static file at the path
	path := r.URL.Path
	if strings.Contains(filename(path), ".") {
		http.ServeFile(w, r, s.pubDir+path)
		return
	}

	// If there's any query at all, serve the static index
	if r.URL.RawQuery != "" {
		http.ServeFile(w, r, s.pubDir+"/index.html")
		return
	}

	// Attempt to render the URL, failing over to the static index
	// if we fail
	body, err := s.renderer.Render(r.URL.String())
	if err != nil {
		http.ServeFile(w, r, s.pubDir+"/index.html")
	}
	w.Write(body)
}

// filename takes a path and returns everything after the the
// final slash in a path
func filename(path string) string {
	// Find the index of the last occuring slash
	i := strings.LastIndex(path, "/")
	if i == -1 {
		return ""
	}

	// The filename is everyting after the final slash
	return path[i+1:]
}

// insContains returns true if any of the substrings are found
// in the passed string. This method of checking contains is
// case insensitive
func insContains(s string, substr string) bool {
	// Always return false if an empty substring
	if substr == "" {
		return false
	}

	// Return whether or not the passed insensitive substring
	// exists in the passed insensitive string
	s, substr = strings.ToLower(s), strings.ToLower(substr)
	return strings.Contains(s, substr)
}
