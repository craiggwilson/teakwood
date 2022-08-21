package flow

import "github.com/charmbracelet/lipgloss"

func DefaultStyles() Styles {
	return Styles{
		Group: lipgloss.NewStyle(),
		Item:  lipgloss.NewStyle(),
	}
}

type Styles struct {
	Group lipgloss.Style
	Item  lipgloss.Style
}
