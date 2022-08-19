package pages

import tea "github.com/charmbracelet/bubbletea"

func New(pages ...tea.Model) Model {
	return Model{
		pages: pages,
	}
}

type Model struct {
	pages []tea.Model

	currentPage int

	width  int
	height int
}

func (m *Model) AddPage(v tea.Model) {
	m.pages = append(m.pages, v)
}

func (m Model) CurrentPage() int {
	return m.currentPage
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) InsertPage(v int, page tea.Model) {
	m.pages = append(m.pages[:v+1], m.pages[v:]...)
	m.pages[v] = page
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

func (m *Model) RemovePage(v int) {
	m.pages = append(m.pages[:v], m.pages[v+1:]...)
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

func (m Model) View() string {
	return m.pages[m.currentPage].View()
}
