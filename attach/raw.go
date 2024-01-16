package attach

// Raw is a raw attach with content in memory
type Raw struct {
	name        string
	contentType string
	content     []byte
}

// NewRaw creates a new raw attach
func NewRaw(name, contentType string, content []byte) *Raw {
	return &Raw{
		name:        name,
		contentType: contentType,
		content:     content,
	}
}

// Name returns the name of the attach
func (a *Raw) Name() string { return a.name }

// ContentType returns the content type of the attach
func (a *Raw) ContentType() string { return a.contentType }

// Content returns the content of the attach
func (a *Raw) Content() []byte { return a.content }
