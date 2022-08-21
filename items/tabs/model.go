package tabs

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teakwood"
	"github.com/craiggwilson/teakwood/items"
)

func New(itemsSource items.Source, opts ...Opt) Model {
	m := Model{
		itemsSource: itemsSource,
		styles:      DefaultStyles(),
	}

	for _, opt := range opts {
		opt(&m)
	}

	return m
}

type Model struct {
	bounds      teakwood.Rectangle
	currentTab  int
	items       []tea.Model
	itemsSource items.Source
	styles      Styles
}

func (m Model) CurrentTab() int {
	return m.currentTab
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) NextTab() {
	if m.currentTab+1 < m.itemsSource.Len() {
		m.currentTab++
	}
}

func (m *Model) PrevTab() {
	if m.currentTab > 0 {
		m.currentTab--
	}
}

func (m *Model) SetCurrentTab(v int) {
	switch {
	case v < 0:
		m.currentTab = 0
	case v >= m.itemsSource.Len():
		m.currentTab = m.itemsSource.Len() - 1
	default:
		m.currentTab = v
	}
}

func (m *Model) SetItemsSource(itemsSource items.Source) {
	m.itemsSource = itemsSource
}

func (m *Model) SetStyles(styles Styles) {
	m.styles = styles
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	m.items = make([]tea.Model, m.itemsSource.Len())
	for i := 0; i < m.itemsSource.Len(); i++ {
		m.items[i], cmd = m.itemsSource.Item(i).Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) UpdateBounds(bounds teakwood.Rectangle) teakwood.Visual {
	m.bounds = bounds
	return m
}

func (m Model) View() string {
	views := make([]string, len(m.items))
	for i := 0; i < len(views); i++ {
		view := m.items[i].View()
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
