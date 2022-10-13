package theme

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teakwood"
)

func New() *Builder {
	return &Builder{
		t: defaultTheme(),
	}
}

type Builder struct {
	t *theme
}

func (b *Builder) Set(key string, style lipgloss.Style) *Builder {
	s, ok := b.t.styles[key]
	if ok {
		style = style.Inherit(s)
	}
	b.t.styles[key] = style
	return b
}

func (b *Builder) Build() teakwood.Styler {
	return b.t
}
