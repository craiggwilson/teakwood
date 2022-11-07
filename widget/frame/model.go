package label

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/craiggwilson/teakwood"
	"github.com/craiggwilson/teakwood/widget"
)

func New(content teakwood.Visual, opts ...Opt) *Model {
	m := Model{
		content: content,
	}

	m.widget = widget.New("label", m.render)
	m.widget.AddChildren(content)
	m.Visual = m.widget

	for _, opt := range opts {
		opt(&m)
	}

	return &m
}

type Model struct {
	teakwood.Visual

	widget  *widget.Widget
	content teakwood.Visual
}

func (m *Model) Content() teakwood.Visual {
	return m.content
}

func (m *Model) SetContent(content teakwood.Visual) {
	m.content = content
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	mdl, cmd := m.Visual.Update(msg)
	m.Visual = mdl.(*widget.Widget)
	return m, cmd
}

func (m *Model) render(style lipgloss.Style) string {
	return style.Render(m.content.View())
}
