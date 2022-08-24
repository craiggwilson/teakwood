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

	bounds teakwood.Rectangle
}

func (m Model) Init() tea.Cmd {
	return m.content.Init()
}

func (m *Model) SetContent(content tea.Model) {
	m.content = content
}

func (m *Model) SetStyle(style lipgloss.Style) {
	m.style = style
	m.applyBoundsToStyleAndContent()
}

func (m Model) Style() lipgloss.Style {
	return m.style
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.content, cmd = m.content.Update(msg)
	return m, cmd
}

func (m Model) UpdateBounds(bounds teakwood.Rectangle) teakwood.Visual {
	m.bounds = bounds
	m.applyBoundsToStyleAndContent()
	return m
}

func (m Model) View() string {
	return m.style.Render(m.content.View())
}

func (m *Model) applyBoundsToStyleAndContent() {
	offsets := teakwood.OffsetsFromStyle(m.style)
	m.style = m.style.Width(m.bounds.Width - offsets.Width).Height(m.bounds.Height - offsets.Height)
	m.content = teakwood.UpdateBounds(m.content, m.bounds.Offset(offsets))
}
