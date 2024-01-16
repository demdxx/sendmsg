package godef

import (
	"context"
	htmltemplate "html/template"
	"io"
	"regexp"
	texttemplate "text/template"

	va "gitgub.com/demdxx/sendmsg/internal/varaccessor"
)

var varReplacer = regexp.MustCompile(`\{\{\s*\.([a-zA-Z0-9_]+)\s*\}\}`)

type itmpl interface {
	Execute(wr io.Writer, data any) error
}

// Template is the template for the message
type Template[TMPL itmpl] struct {
	name    string
	isHTML  bool
	subject string
	tmpl    itmpl
}

type (
	HTMLTemplate = Template[*htmltemplate.Template]
	TextTemplate = Template[*texttemplate.Template]
)

// NewHTMLTemplate returns new template
func NewHTMLTemplate(name, subject, content string) (*HTMLTemplate, error) {
	tmpl, err := htmltemplate.New("default").Parse(content)
	if err != nil {
		return nil, err
	}
	return &HTMLTemplate{name: name, subject: subject, tmpl: tmpl, isHTML: true}, nil
}

// NewTextTemplate returns new template
func NewTextTemplate(name, subject, content string) (*TextTemplate, error) {
	tmpl, err := texttemplate.New("default").Parse(content)
	if err != nil {
		return nil, err
	}
	return &TextTemplate{name: name, subject: subject, tmpl: tmpl, isHTML: false}, nil
}

// Name returns the name of the template
func (t *Template[TMPL]) Name() string {
	return t.name
}

// IsHTML returns true if the template is HTML
func (t *Template[TMPL]) IsHTML() bool {
	return t.isHTML
}

// Subject returns the subject of the template
func (t *Template[TMPL]) Subject(ctx context.Context, vars map[string]any) (string, error) {
	return va.ReplaceVarsWithRe(t.subject, vars, varReplacer), nil
}

// Render renders the template with the given params
func (t *Template[TMPL]) Render(ctx context.Context, wr io.Writer, params map[string]any) error {
	if err := t.tmpl.Execute(wr, params); err != nil {
		return err
	}
	return nil
}
