package sendmsg

import (
	"bytes"
	"context"
	"io"

	"github.com/demdxx/xtypes"
)

type itemplater interface {
	IsHTML() bool
	Subject(ctx context.Context, vars map[string]any) (string, error)
	Render(ctx context.Context, wr io.Writer, vars map[string]any) error
}

type TemplateMessage struct {
	Recipients []string
	CC         []string
	BCC        []string
	Template   itemplater
	Attaches   []Attach
	Vars       map[string]any
}

func (m *TemplateMessage) GetRecipients() []string { return m.Recipients }
func (m *TemplateMessage) GetCC() []string         { return m.CC }
func (m *TemplateMessage) GetBCC() []string        { return m.BCC }
func (m *TemplateMessage) GetAttaches() []Attach   { return m.Attaches }

func (m *TemplateMessage) GetSubject(ctx context.Context, vars map[string]any) (string, error) {
	return m.Template.Subject(ctx, xtypes.Map[string, any](m.Vars).Merge(vars))
}

func (m *TemplateMessage) GetHTML(ctx context.Context, vars map[string]any) (string, error) {
	if m.Template != nil && m.Template.IsHTML() {
		return m.renderTmpl(ctx, vars)
	}
	return "", nil
}

func (m *TemplateMessage) GetPlainText(ctx context.Context, vars map[string]any) (string, error) {
	if m.Template != nil && !m.Template.IsHTML() {
		return m.renderTmpl(ctx, vars)
	}
	return "", nil
}

func (m *TemplateMessage) renderTmpl(ctx context.Context, vars map[string]any) (string, error) {
	var buf bytes.Buffer
	err := m.Template.Render(ctx, &buf, xtypes.Map[string, any](m.Vars).Merge(vars))
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (m *TemplateMessage) Complete() error { return nil }

var _ Message = (*TemplateMessage)(nil)
