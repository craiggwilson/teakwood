package tabs

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teakwood"
)

func New(tabs ...tea.Model) Model {
	return Model{
		tabs:   tabs,
		styles: DefaultStyles(),
	}
}

type Model struct {
	tabs []tea.Model

	currentTab int

	styles Styles

	bounds teakwood.Rectangle
}

func (m *Model) AddTab(v tea.Model) {
	m.tabs = append(m.tabs, v)
}

func (m Model) CurrentTab() int {
	return m.currentTab
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) InsertTab(v int, tab tea.Model) {
	m.tabs = append(m.tabs[:v+1], m.tabs[v:]...)
	m.tabs[v] = tab
}

func (m *Model) NextTab() {
	if m.currentTab+1 < len(m.tabs) {
		m.currentTab++
	}
}

func (m *Model) PrevTab() {
	if m.currentTab > 0 {
		m.currentTab--
	}
}

func (m *Model) RemoveTab(v int) {
	m.tabs = append(m.tabs[:v], m.tabs[v+1:]...)
}

func (m *Model) SetCurrentTab(v int) {
	switch {
	case v < 0:
		m.currentTab = 0
	case v >= len(m.tabs):
		m.currentTab = len(m.tabs) - 1
	default:
		m.currentTab = v
	}
}

func (m *Model) SetStyles(styles Styles) {
	m.styles = styles
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	for i, tab := range m.tabs {
		m.tabs[i], cmd = tab.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) UpdateBounds(bounds teakwood.Rectangle) teakwood.Visual {
	m.bounds = bounds
	return m
}

func (m Model) View() string {
	views := make([]string, len(m.tabs))
	for i, tab := range m.tabs {
		view := tab.View()
		if m.currentTab == i {
			view = m.styles.SelectedTab.Render(view)
		} else {
			view = m.styles.Tab.Render(view)
		}

		views[i] = view
	}

	view := lipgloss.JoinHorizontal(lipgloss.Top, views...)
	filler := m.styles.Filler.Render(strings.Repeat(" ", max(0, m.bounds.Width-lipgloss.Width(view))))
	return lipgloss.JoinHorizontal(lipgloss.Bottom, view, filler)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
