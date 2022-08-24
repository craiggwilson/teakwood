package teakwood

import tea "github.com/charmbracelet/bubbletea"

type Visual interface {
	tea.Model

	UpdateBounds(Rectangle) Visual
}

func UpdateBounds(m tea.Model, bounds Rectangle) tea.Model {
	if v, ok := any(m).(Visual); ok {
		return v.UpdateBounds(bounds)
	}

	return m
}
