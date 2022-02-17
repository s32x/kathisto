package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilename(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("file", filename("localhost:8080/test/test/file"))
	assert.Equal("", filename("localhost:8080"))
	assert.Equal("", filename("localhost:8080/"))
	assert.Equal("", filename("localhost:8080/test/test/"))
	assert.Equal("robots.txt", filename("localhost:8080/test/test/robots.txt"))
	assert.Equal("sitemap.xml", filename("localhost:8080/test/test/sitemap.xml"))
}

func TestInscontains(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, insContains("example.com", "EXAMPLE.com"))
	assert.Equal(false, insContains("example.com", ""))
	assert.Equal(true, insContains("example.com", "example.com"))
}
