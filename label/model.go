package label

import (
	tea "github.com/charmbracelet/bubbletea"
)

func New(text string) Model {
	return Model(text)
}

type Model string

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return string(m)
}
