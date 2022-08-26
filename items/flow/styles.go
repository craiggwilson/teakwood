package flow

import "github.com/charmbracelet/lipgloss"

func DefaultStyles() Styles {
	return Styles{
		Group:        lipgloss.NewStyle(),
		Item:         lipgloss.NewStyle(),
		SelectedItem: lipgloss.NewStyle(),
	}
}

type Styles struct {
	Group        lipgloss.Style
	Item         lipgloss.Style
	SelectedItem lipgloss.Style
	CurrentItem  lipgloss.Style
}
