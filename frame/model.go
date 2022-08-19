package frame

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teacomps/sizeutil"
)

func New(style lipgloss.Style, child tea.Model) Model {
	return Model{
		child: child,
		style: style,
	}
}

type Model struct {
	child tea.Model
	style lipgloss.Style
}

func (m Model) Init() tea.Cmd {
	return m.child.Init()
}

func (m *Model) SetHeight(v int) {
	margins := m.style.GetVerticalFrameSize()
	m.style = m.style.Height(v - margins)
	m.child, _ = sizeutil.TrySetHeight(m.child, v-margins)
}

func (m *Model) SetStyle(v lipgloss.Style) {
	m.style = v
}

func (m *Model) SetWidth(v int) {
	margins := m.style.GetHorizontalFrameSize()
	m.style = m.style.Width(v - margins)
	m.child, _ = sizeutil.TrySetWidth(m.child, v-margins)
}

func (m Model) Style() lipgloss.Style {
	return m.style
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.child, cmd = m.child.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.style.Render(m.child.View())
}
