package attach

import "io"

// IO is a io attach
type IO struct {
	name        string
	contentType string
	content     io.Reader
}

// NewIO creates a new io attach
func NewIO(name, contentType string, content io.Reader) *IO {
	return &IO{
		name:        name,
		contentType: contentType,
		content:     content,
	}
}

// Name returns the name of the attach
func (a *IO) Name() string { return a.name }

// ContentType returns the content type of the attach
func (a *IO) ContentType() string { return a.contentType }

// Content returns the content of the attach
func (a *IO) Content() io.Reader { return a.content }

// Close closes the content reader
func (a *IO) Close() error {
	if closer, ok := a.content.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}
