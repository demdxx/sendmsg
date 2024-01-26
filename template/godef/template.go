package godef

import (
	"context"
	htmltemplate "html/template"
	"io"
	"regexp"
	texttemplate "text/template"

	va "github.com/demdxx/sendmsg/internal/varaccessor"
	"github.com/demdxx/xtypes"
)

type preRenderFunc func(ctx context.Context, vars map[string]any) (map[string]any, error)

// varReplacer is the regexp for replace vars: {{ .var.name.subname }}
var varReplacer = regexp.MustCompile(`\{\{\s*\.([a-zA-Z0-9_]+)\s*\}\}`)

type itmpl interface {
	Execute(wr io.Writer, data any) error
}

// Template is the template for the message
type Template[TMPL itmpl] struct {
	name      string
	isHTML    bool
	subject   string
	tmpl      itmpl
	vars      map[string]any
	preRender preRenderFunc
}

type (
	HTMLTemplate = Template[*htmltemplate.Template]
	TextTemplate = Template[*texttemplate.Template]
)

// NewHTMLTemplate returns new template
func NewHTMLTemplate(name, subject, content string, opts ...Option[*htmltemplate.Template]) (*HTMLTemplate, error) {
	tmpl, err := htmltemplate.New("default").Parse(content)
	if err != nil {
		return nil, err
	}
	tmplObj := &HTMLTemplate{
		name:    name,
		subject: subject,
		tmpl:    tmpl,
		isHTML:  true,
	}
	for _, opt := range opts {
		opt(tmplObj)
	}
	return tmplObj, nil
}

// NewTextTemplate returns new template
func NewTextTemplate(name, subject, content string, opts ...Option[*texttemplate.Template]) (*TextTemplate, error) {
	tmpl, err := texttemplate.New("default").Parse(content)
	if err != nil {
		return nil, err
	}
	tmplObj := &TextTemplate{
		name:    name,
		subject: subject,
		tmpl:    tmpl,
		isHTML:  false,
	}
	for _, opt := range opts {
		opt(tmplObj)
	}
	return tmplObj, nil
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
	extVars := xtypes.Map[string, any](t.vars).Merge(vars, t.vars)
	if t.preRender != nil {
		var err error
		if extVars, err = t.preRender(ctx, extVars); err != nil {
			return "", err
		}
	}
	return va.ReplaceVarsWithRe(t.subject, extVars, varReplacer), nil
}

// Render renders the template with the given params
func (t *Template[TMPL]) Render(ctx context.Context, wr io.Writer, params map[string]any) error {
	extParams := xtypes.Map[string, any](t.vars).Merge(params, t.vars)
	if t.preRender != nil {
		var err error
		if extParams, err = t.preRender(ctx, extParams); err != nil {
			return err
		}
	}
	if err := t.tmpl.Execute(wr, extParams); err != nil {
		return err
	}
	return nil
}
