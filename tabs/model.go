package tabs

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teakwood"
)

func New(opts ...Opt) Model {
	m := Model{
		styles: DefaultStyles(),
	}

	for _, opt := range opts {
		opt(&m)
	}

	return m
}

type Model struct {
	bounds     teakwood.Rectangle
	currentTab int
	items      []tea.Model
	styles     Styles

	itemViews []string
}

func (m *Model) AddItem(tab tea.Model) {
	m.items = append(m.items, tab)
}

func (m Model) CurrentTab() int {
	return m.currentTab
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) InsertItem(index int, page tea.Model) {
	m.items = append(m.items[:index+1], m.items[index:]...)
	m.items[index] = page
}

func (m Model) Items() []tea.Model {
	return m.items
}

func (m *Model) NextTab() {
	if m.currentTab+1 < len(m.items) {
		m.currentTab++
	}
}

func (m *Model) PrevTab() {
	if m.currentTab > 0 {
		m.currentTab--
	}
}

func (m *Model) RemoveItem(index int) {
	m.items = append(m.items[:index], m.items[index+1:]...)
}

func (m *Model) SetCurrentItem(v int) {
	switch {
	case v < 0:
		m.currentTab = 0
	case v >= len(m.items):
		m.currentTab = len(m.items) - 1
	default:
		m.currentTab = v
	}
}

func (m *Model) SetItems(items ...tea.Model) {
	m.items = items
}

func (m *Model) SetStyles(styles Styles) {
	m.styles = styles
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	for i := 0; i < len(m.items); i++ {
		m.items[i], cmd = m.items[i].Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) UpdateBounds(bounds teakwood.Rectangle) teakwood.Visual {
	m.bounds = bounds
	return m
}

func (m Model) View() string {
	if len(m.itemViews) != len(m.items) {
		m.itemViews = make([]string, len(m.items))
	}

	for i := range m.items {

		view := m.items[i].View()
		if m.currentTab == i {
			view = m.styles.SelectedTab.Render(view)
		} else {
			view = m.styles.Tab.Render(view)
		}

		m.itemViews[i] = view
	}

	itemsView := lipgloss.JoinHorizontal(lipgloss.Top, m.itemViews...)
	filler := m.styles.Filler.Render(strings.Repeat(" ", max(0, m.bounds.Width-lipgloss.Width(itemsView))))
	return lipgloss.JoinHorizontal(lipgloss.Bottom, itemsView, filler)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
