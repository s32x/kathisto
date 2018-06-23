package renderer

import (
	"bytes"
	"fmt"
	"os/exec"
)

// noRenderKey is used in the queryparam of an forced
// un-rendered request to avoid recursive rendering
const noRenderKey = "norender"

// Chrome is a struct containing all dependencies needed
// to perform HTML rendering and response caching using
// Headless Chrome. It satisfies the Renderer interface
type Chrome struct {
	userAgent string
	cache     map[string][]byte
}

// NewChromeRenderer generates a new Chrome Renderer that
// renders and stores rendered pages in a cache
func NewChromeRenderer(userAgent string) Renderer {
	return &Chrome{userAgent, make(map[string][]byte)}
}

// Render renders a page using Headless Chrome and returns
// the bytes
func (c *Chrome) Render(url string) ([]byte, error) {
	// Return any previously rendered cached bytes
	if c.cache[url] != nil {
		return c.cache[url], nil
	}

	// Construct a norender url to avoid recursive
	// rendering
	url = fmt.Sprintf("%s?%s=true", url, noRenderKey)

	// Execute the render command using the norender url
	page, err := exec.Command(
		"chromium-browser",
		"--headless",
		"--no-sandbox",
		"--dump-dom",
		"--user-agent="+c.userAgent,
		url).Output()
	if err != nil {
		return nil, err
	}
	page = stripErrors(page)

	// Cache and return the bytes that have been rendered
	c.cache[url] = page
	return page, nil
}

// stripErrors strips all prefixing errors from the
// rendered webpage
func stripErrors(body []byte) []byte {
	buf := bytes.NewBuffer(body)
	for {
		switch c := string(buf.Next(1)); c {
		case "[":
			buf.ReadBytes('\n')
		default:
			return append([]byte(c), buf.Bytes()...)
		}
	}
}
