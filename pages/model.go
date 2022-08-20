package pages

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/craiggwilson/teakwood"
)

func New(pages ...tea.Model) Model {
	return Model{
		pages: pages,
	}
}

type Model struct {
	pages []tea.Model

	currentPage int

	bounds teakwood.Rectangle
}

func (m *Model) AddPage(page tea.Model) {
	m.pages = append(m.pages, page)
}

func (m Model) CurrentPage() int {
	return m.currentPage
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) InsertPage(index int, page tea.Model) {
	m.pages = append(m.pages[:index+1], m.pages[index:]...)
	m.pages[index] = page
}

func (m *Model) NextPage() {
	if m.currentPage+1 < len(m.pages) {
		m.currentPage++
	}
}

func (m *Model) PrevPage() {
	if m.currentPage > 0 {
		m.currentPage--
	}
}

func (m *Model) RemovePage(index int) {
	m.pages = append(m.pages[:index], m.pages[index+1:]...)
}

func (m *Model) SetCurrentPage(v int) {
	switch {
	case v < 0:
		m.currentPage = 0
	case v >= len(m.pages):
		m.currentPage = len(m.pages) - 1
	default:
		m.currentPage = v
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	for i, page := range m.pages {
		m.pages[i], cmd = page.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) UpdateBounds(bounds teakwood.Rectangle) teakwood.Visual {
	m.bounds = bounds
	for i, page := range m.pages {
		if c, ok := page.(teakwood.Visual); ok {
			m.pages[i] = c.UpdateBounds(bounds)
		}
	}

	return m
}

func (m Model) View() string {
	return m.pages[m.currentPage].View()
}
