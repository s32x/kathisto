package renderer

// Renderer is used for fully rendering html pages on the server
type Renderer interface {
	Render(u string) ([]byte, error)
}
