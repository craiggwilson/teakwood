package frame

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teakwood"
)

func New(style lipgloss.Style, content tea.Model) Model {
	return Model{
		content: content,
		style:   style,
	}
}

type Model struct {
	content tea.Model
	style   lipgloss.Style

	bounds teakwood.Rectangle
}

func (m Model) Init() tea.Cmd {
	return m.content.Init()
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
	topFrame := m.style.GetMarginTop() + m.style.GetPaddingTop() + m.style.GetBorderTopWidth()
	bottomFrame := m.style.GetMarginBottom() + m.style.GetPaddingBottom() + m.style.GetBorderBottomSize()
	leftFrame := m.style.GetMarginLeft() + m.style.GetPaddingLeft() + m.style.GetBorderLeftSize()
	rightFrame := m.style.GetMarginRight() + m.style.GetPaddingRight() + m.style.GetBorderRightSize()

	m.style = m.style.Width(m.bounds.Width - leftFrame - rightFrame).Height(m.bounds.Height - topFrame - bottomFrame)
	if c, ok := m.content.(teakwood.Visual); ok {
		m.content = c.UpdateBounds(teakwood.NewRectangle(
			m.bounds.X+leftFrame,
			m.bounds.Y+topFrame,
			m.bounds.Width-leftFrame-rightFrame,
			m.bounds.Height-topFrame-bottomFrame,
		))
	}
}
