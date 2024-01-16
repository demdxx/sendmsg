package varaccessor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplacer(t *testing.T) {
	text := ReplaceVars("test : {{a}} / {{ b}} / {{ c }}", map[string]any{"a": 1, "b": 2, "c": 3})
	assert.Equal(t, "test : 1 / 2 / 3", text)
}
