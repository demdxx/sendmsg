package godef

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplate(t *testing.T) {
	ctx := context.Background()
	preRenderTest := func(ctx context.Context, vars map[string]any) (map[string]any, error) {
		vars["v1"] = "1"
		return vars, nil
	}

	t.Run("html", func(t *testing.T) {
		tmpl1, err := NewHTMLTemplate("test",
			"test: {{ .title }}{{ .v1 }}", "<h1>{{ .text }}{{ .v1 }}</h1>",
			WithHTMLVars(nil),
			WithHTMLPreRender(preRenderTest))
		if assert.NoError(t, err) {
			var buf bytes.Buffer
			assert.Equal(t, "test", tmpl1.Name())
			if subject, err := tmpl1.Subject(ctx, map[string]any{"title": "title"}); assert.NoError(t, err) {
				assert.Equal(t, "test: title1", subject)
			}
			assert.True(t, tmpl1.IsHTML())
			err = tmpl1.Render(ctx, &buf, map[string]any{"text": "test"})
			assert.NoError(t, err)
			assert.Equal(t, "<h1>test1</h1>", buf.String())
		}
	})

	t.Run("text", func(t *testing.T) {
		tmpl2, err := NewTextTemplate("test",
			"test: {{ .title }}{{ .v1 }}", "text: {{ .text }}{{ .v1 }}",
			WithTextVars(nil),
			WithTextPreRender(preRenderTest))
		if assert.NoError(t, err) {
			var buf bytes.Buffer
			assert.Equal(t, "test", tmpl2.Name())
			if subject, err := tmpl2.Subject(ctx, map[string]any{"title": "title"}); assert.NoError(t, err) {
				assert.Equal(t, "test: title1", subject)
			}
			assert.False(t, tmpl2.IsHTML())
			err = tmpl2.Render(ctx, &buf, map[string]any{"text": "test"})
			assert.NoError(t, err)
			assert.Equal(t, "text: test1", buf.String())
		}
	})
}
