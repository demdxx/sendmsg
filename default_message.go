package sendmsg

import (
	"context"

	"github.com/demdxx/xtypes"

	va "gitgub.com/demdxx/sendmsg/internal/varaccessor"
)

type DefaultMessage struct {
	Recipients []string
	CC         []string
	BCC        []string
	Subject    string
	HTML       string
	PlainText  string
	Attaches   []Attach
	Vars       map[string]any
}

func (m *DefaultMessage) GetRecipients() []string { return m.Recipients }
func (m *DefaultMessage) GetCC() []string         { return m.CC }
func (m *DefaultMessage) GetBCC() []string        { return m.BCC }
func (m *DefaultMessage) GetAttaches() []Attach   { return m.Attaches }

func (m *DefaultMessage) GetSubject(ctx context.Context, vars map[string]any) (string, error) {
	return va.ReplaceVars(m.Subject, xtypes.Map[string, any](m.Vars).Merge(vars)), nil
}

func (m *DefaultMessage) GetHTML(ctx context.Context, vars map[string]any) (string, error) {
	return va.ReplaceVars(m.HTML, xtypes.Map[string, any](m.Vars).Merge(vars)), nil
}

func (m *DefaultMessage) GetPlainText(ctx context.Context, vars map[string]any) (string, error) {
	return va.ReplaceVars(m.PlainText, xtypes.Map[string, any](m.Vars).Merge(vars)), nil
}

func (m *DefaultMessage) Complete() error { return nil }

var _ Message = (*DefaultMessage)(nil)
