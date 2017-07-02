package liquid

import (
	"io"

	"github.com/osteele/liquid/chunks"
	"github.com/osteele/liquid/filters"
	"github.com/osteele/liquid/tags"
)

type engine struct{ settings chunks.Settings }

// NewEngine returns a new template engine.
func NewEngine() Engine {
	e := engine{chunks.NewSettings()}
	filters.AddStandardFilters(e.settings.ExpressionSettings)
	tags.AddStandardTags(e.settings)
	return e
}

// RegisterBlock is in the Engine interface.
func (e engine) RegisterBlock(name string, td func(io.Writer, chunks.RenderContext) error) {
	e.settings.AddBlock(name).Renderer(td)
}

// RegisterFilter is in the Engine interface.
func (e engine) RegisterFilter(name string, fn interface{}) {
	e.settings.AddFilter(name, fn)
}

// RegisterTag is in the Engine interface.
func (e engine) RegisterTag(name string, td TagDefinition) {
	// For simplicity, don't expose the two stage parsing/rendering process to clients.
	// Client tags do everything at runtime.
	e.settings.AddTag(name, func(_ string) (func(io.Writer, chunks.RenderContext) error, error) {
		return td, nil
	})
}

// ParseTemplate is in the Engine interface.
func (e engine) ParseTemplate(text []byte) (Template, error) {
	ast, err := e.settings.Parse(string(text))
	if err != nil {
		return nil, err
	}
	return &template{ast, e.settings}, nil
}

// ParseAndRender is in the Engine interface.
func (e engine) ParseAndRender(text []byte, c Context) ([]byte, error) {
	t, err := e.ParseTemplate(text)
	if err != nil {
		return nil, err
	}
	return t.Render(c)
}

// ParseAndRenderString is in the Engine interface.
func (e engine) ParseAndRenderString(text string, c Context) (string, error) {
	b, err := e.ParseAndRender([]byte(text), c)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
