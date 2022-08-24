package pages

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/craiggwilson/teakwood"
)

func New(items ...tea.Model) Model {
	return Model{
		items: items,
	}
}

type Model struct {
	items []tea.Model

	currentPage int

	bounds teakwood.Rectangle
}

func (m *Model) AddItem(page tea.Model) {
	m.items = append(m.items, page)
}

func (m Model) CurrentPage() int {
	return m.currentPage
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

func (m Model) Len() int {
	return len(m.items)
}

func (m *Model) NextPage() {
	if m.currentPage+1 < len(m.items) {
		m.currentPage++
	}
}

func (m *Model) PrevPage() {
	if m.currentPage > 0 {
		m.currentPage--
	}
}

func (m *Model) RemoveItem(index int) {
	m.items = append(m.items[:index], m.items[index+1:]...)
}

func (m *Model) SetCurrentItem(v int) {
	switch {
	case v < 0:
		m.currentPage = 0
	case v >= len(m.items):
		m.currentPage = len(m.items) - 1
	default:
		m.currentPage = v
	}
}

func (m *Model) SetItems(items ...tea.Model) {
	m.items = items
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	for i, page := range m.items {
		m.items[i], cmd = page.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) UpdateBounds(bounds teakwood.Rectangle) teakwood.Visual {
	m.bounds = bounds
	for i := range m.items {
		m.items[i] = teakwood.UpdateBounds(m.items[i], m.bounds)
	}

	return m
}

func (m Model) View() string {
	return m.items[m.currentPage].View()
}
