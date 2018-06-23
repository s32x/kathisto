package renderer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	errRender = `[error]test-error-text
[error2]test-error-test
<html></html>`

	realError = `[0623/085520.358025:ERROR:gpu_process_transport_factory.cc(1007)] Lost UI shared context.
[0623/085520.358025:ERROR:gpu_process_transport_factory.cc(1007)] Lost UI shared context.
<html></html>`

	successRender = `<html></html>`
)

func TestStripErrors(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("<html></html>", string(stripErrors([]byte(errRender))))
	assert.Equal("<html></html>", string(stripErrors([]byte(realError))))
	assert.Equal("<html></html>", string(stripErrors([]byte(successRender))))
}
