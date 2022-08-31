package tabs

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teakwood"
	"github.com/craiggwilson/teakwood/items"
)

func New[T any](source items.Source[T], renderer items.Renderer[T], opts ...Opt[T]) Model[T] {
	m := Model[T]{
		renderer: renderer,
		source:   source,
		styles:   DefaultStyles(),
	}

	for _, opt := range opts {
		opt(&m)
	}

	return m
}

type Model[T any] struct {
	currentTab int
	renderer   items.Renderer[T]
	source     items.Source[T]
	styles     Styles

	views []string
}

func (m Model[T]) CurrentTab() int {
	return m.currentTab
}

func (m Model[T]) Init() tea.Cmd {
	return nil
}

func (m *Model[T]) NextTab() {
	if m.currentTab+1 < m.source.Len() {
		m.currentTab++
	}
}

func (m *Model[T]) PrevTab() {
	if m.currentTab > 0 {
		m.currentTab--
	}
}

func (m Model[T]) Renderer() items.Renderer[T] {
	return m.renderer
}

func (m *Model[T]) SetCurrentItem(v int) {
	switch {
	case v < 0:
		m.currentTab = 0
	case v >= m.source.Len():
		m.currentTab = m.source.Len() - 1
	default:
		m.currentTab = v
	}
}

func (m *Model[T]) SetRenderer(renderer items.Renderer[T]) {
	m.renderer = renderer
}

func (m *Model[T]) SetSource(source items.Source[T]) {
	m.source = source
}

func (m *Model[T]) SetStyles(styles Styles) {
	m.styles = styles
}

func (m Model[T]) Source() items.Source[T] {
	return m.source
}

func (m Model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model[T]) View() string {
	return m.renderTabs()
}

func (m Model[T]) ViewWithBounds(bounds teakwood.Rectangle) string {
	itemsView := m.renderTabs()
	filler := m.styles.Filler.Render(strings.Repeat(" ", max(0, bounds.Width-lipgloss.Width(itemsView))))
	return lipgloss.JoinHorizontal(lipgloss.Bottom, itemsView, filler)
}

func (m Model[T]) renderTabs() string {
	if len(m.views) != m.source.Len() {
		m.views = make([]string, m.source.Len())
	}

	for i := 0; i < m.source.Len(); i++ {

		view := m.renderer.Render(m.source.Item(i))
		if m.currentTab == i {
			view = m.styles.SelectedItem.Render(view)
		} else {
			view = m.styles.Item.Render(view)
		}

		m.views[i] = view
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, m.views...)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
