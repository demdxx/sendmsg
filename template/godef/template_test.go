package godef

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplate(t *testing.T) {
	ctx := context.Background()

	t.Run("html", func(t *testing.T) {
		tmpl1, err := NewHTMLTemplate("test", "test: {{ .title }}", "<h1>{{ .text }}</h1>")
		if assert.NoError(t, err) {
			var buf bytes.Buffer
			assert.Equal(t, "test", tmpl1.Name())
			if subject, err := tmpl1.Subject(ctx, map[string]any{"title": "title"}); assert.NoError(t, err) {
				assert.Equal(t, "test: title", subject)
			}
			assert.True(t, tmpl1.IsHTML())
			err = tmpl1.Render(ctx, &buf, map[string]any{"text": "test"})
			assert.NoError(t, err)
			assert.Equal(t, "<h1>test</h1>", buf.String())
		}
	})

	t.Run("text", func(t *testing.T) {
		tmpl2, err := NewTextTemplate("test", "test: {{ .title }}", "text: {{ .text }}")
		if assert.NoError(t, err) {
			var buf bytes.Buffer
			assert.Equal(t, "test", tmpl2.Name())
			if subject, err := tmpl2.Subject(ctx, map[string]any{"title": "title"}); assert.NoError(t, err) {
				assert.Equal(t, "test: title", subject)
			}
			assert.False(t, tmpl2.IsHTML())
			err = tmpl2.Render(ctx, &buf, map[string]any{"text": "test"})
			assert.NoError(t, err)
			assert.Equal(t, "text: test", buf.String())
		}
	})
}
