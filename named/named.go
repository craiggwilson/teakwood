package named

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/craiggwilson/teakwood"
)

func New(name string, content tea.Model) Model {
	return Model{name: name, content: content}
}

type Model struct {
	name    string
	content tea.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Name() string {
	return m.name
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch tmsg := msg.(type) {
	case UpdateMsg:
		if tmsg.Name == m.name {
			m.content, cmd = tmsg.Update(m.content, msg)
			cmds = append(cmds, cmd)
		}
	}

	m.content, cmd = m.content.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) UpdateBounds(bounds teakwood.Rectangle) teakwood.Visual {
	if c, ok := m.content.(teakwood.Visual); ok {
		m.content = c.UpdateBounds(bounds)
	}

	return m
}

func (m Model) View() string {
	return m.content.View()
}

type UpdateFunc func(tea.Model, tea.Msg) (tea.Model, tea.Cmd)

func Update[T any](name string, updater func(T, tea.Msg) (tea.Model, tea.Cmd)) tea.Cmd {
	return func() tea.Msg {
		return UpdateMsg{Name: name, Update: func(m tea.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
			return updater(m.(T), msg)
		}}
	}
}

type UpdateMsg struct {
	Name   string
	Update UpdateFunc
}
