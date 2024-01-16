package varaccessor

import (
	"strings"

	"github.com/demdxx/gocast/v2"
)

// Vars accessor
type Vars struct {
	data map[string]any
}

// NewForAny creates new Vars accessor from any data
func NewForAny(data any) (*Vars, error) {
	m, err := gocast.TryMap[string, any](data)
	if err != nil {
		return nil, err
	}
	return New(m), nil
}

// New creates new Vars accessor
func New(data map[string]any) *Vars {
	return &Vars{data: data}
}

// TryGet returns variable value by key
func (v *Vars) TryGet(key string) (any, error) {
	if !strings.Contains(key, ".") {
		return v.data[key], nil
	}
	keys := strings.Split(key, ".")
	root := v.data[keys[0]]
	if root == nil {
		return nil, nil
	}
	for _, key := range keys[1:] {
		switch {
		case gocast.IsSlice(root):
			idx, err := gocast.TryNumber[uint](key)
			if err != nil {
				return nil, err
			}
			sl, err := gocast.TryAnySlice[any](root)
			if err != nil {
				return nil, err
			}
			if idx < uint(len(sl)) {
				root = sl[idx]
			} else {
				root = nil
			}
		default:
			m, err := gocast.TryMap[string, any](root)
			if err != nil {
				return nil, err
			}
			root = m[key]
		}
		if root == nil {
			return nil, nil
		}
	}
	return root, nil
}

// Get returns variable value by key or default value
func (v *Vars) Get(key string, def ...any) any {
	val, _ := v.TryGet(key)
	if val == nil && len(def) > 0 {
		return def[0]
	}
	return val
}

// String returns variable value by key or default value
func (v *Vars) String(key string, def ...any) string {
	return gocast.Str(v.Get(key, def...))
}

// Int returns variable value by key or default value
func (v *Vars) Int(key string, def ...any) int {
	return gocast.Int(v.Get(key, def...))
}

// Float64 returns variable value by key or default value
func (v *Vars) Float64(key string, def ...any) float64 {
	return gocast.Float64(v.Get(key, def...))
}
