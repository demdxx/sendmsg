package godef

import (
	htmltemplate "html/template"
	texttemplate "text/template"
)

type Option[TMPL itmpl] func(*Template[TMPL])

func WithVars[TMPL itmpl](vars map[string]any) Option[TMPL] {
	return func(t *Template[TMPL]) {
		t.vars = vars
	}
}

func WithTextVars(vars map[string]any) Option[*texttemplate.Template] {
	return WithVars[*texttemplate.Template](vars)
}

func WithHTMLVars(vars map[string]any) Option[*htmltemplate.Template] {
	return WithVars[*htmltemplate.Template](vars)
}

func WithPreRender[TMPL itmpl](preRender preRenderFunc) Option[TMPL] {
	return func(t *Template[TMPL]) {
		t.preRender = preRender
	}
}

func WithTextPreRender(preRender preRenderFunc) Option[*texttemplate.Template] {
	return WithPreRender[*texttemplate.Template](preRender)
}

func WithHTMLPreRender(preRender preRenderFunc) Option[*htmltemplate.Template] {
	return WithPreRender[*htmltemplate.Template](preRender)
}
