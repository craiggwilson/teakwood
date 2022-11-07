package label

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/craiggwilson/teakwood"
	"github.com/craiggwilson/teakwood/widget"
)

func New(text string, opts ...Opt) *Model {
	m := Model{
		text: text,
	}

	m.Visual = widget.New("label", m.render)

	for _, opt := range opts {
		opt(&m)
	}

	return &m
}

type Model struct {
	teakwood.Visual

	text string
}

func (m *Model) SetText(text string) {
	m.text = text
}

func (m *Model) Text() string {
	return m.text
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	mdl, cmd := m.Visual.Update(msg)
	m.Visual = mdl.(*widget.Widget)
	return m, cmd
}

func (m *Model) render(style lipgloss.Style) string {
	return style.Render(m.text)
}
