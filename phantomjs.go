package main

import (
	"fmt"
	"math/rand"
	"net/url"
	"os/exec"
	"time"

	"github.com/patrickmn/go-cache"
)

// Renderer is used for fully rendering html pages on the server
type Renderer interface {
	Render(u string) ([]byte, error)
	IsNoOp(q url.Values) bool // Needed to avoid infinite request loop
}

// phantomJS is a struct containing all dependencies needed to perform
// HTML rendering and periodic response caching using PhantomJS
type phantomJS struct {
	cache              *cache.Cache
	noopKey, userAgent string
}

// NewPJSRenderer generates a new PhantomJS Renderer that contains an
// in-memory cache that expires at the passed expiration
func NewPJSRenderer(expiration time.Duration, userAgent string) Renderer {
	return &phantomJS{noopKey: randSeq(10), userAgent: userAgent}
}

// Render renders a page using PhantomJS and the script we provide it with
func (p *phantomJS) Render(url string) ([]byte, error) {
	// Assemble our PhantomJS command argument (a url that we want to return to
	// the user in a pre-rendered fashion) and executes it on our PhantomJS open script
	noopURL := fmt.Sprintf("%s?%s=%s", url, p.noopKey, "true")
	page, err := exec.Command("phantomjs", "render.js", noopURL, p.userAgent).Output()
	if err != nil {
		return nil, err
	}
	return page, nil
}

// isNoOp is a method used to allow the Renderer to request a resource without
// the request recursively reaching back and hitting the Renderer again
func (p *phantomJS) IsNoOp(query url.Values) bool {
	if query.Get(p.noopKey) == "true" {
		return true
	}
	return false
}

// randSeq generates a random sequence of runes of fixed length n. This is used
// to generate a unique query key for NoOp operations
func randSeq(n int) string {
	rand.Seed(time.Now().UnixNano())
	var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}
