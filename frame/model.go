package frame

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teakwood"
)

func New(content tea.Model, opts ...Opt) Model {
	m := Model{
		content: content,
	}

	for _, opt := range opts {
		opt(&m)
	}

	return m
}

type Model struct {
	content tea.Model
	style   lipgloss.Style
}

func (m Model) Init() tea.Cmd {
	return m.content.Init()
}

func (m *Model) SetContent(content tea.Model) {
	m.content = content
}

func (m *Model) SetStyle(style lipgloss.Style) {
	m.style = style
}

func (m Model) Style() lipgloss.Style {
	return m.style
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.content, cmd = m.content.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.style.Render(m.content.View())
}

func (m Model) ViewWithBounds(bounds teakwood.Rectangle) string {
	offsets := teakwood.OffsetsFromStyle(m.style)

	s := m.style.Copy().Width(bounds.Width - offsets.Width).Height(bounds.Height - offsets.Height)

	if v, ok := m.content.(teakwood.Visual); ok {
		return s.Render(v.ViewWithBounds(bounds.Offset(offsets)))
	}

	return s.Render(m.content.View())
}
