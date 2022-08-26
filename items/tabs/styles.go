package tabs

import "github.com/charmbracelet/lipgloss"

func DefaultStyles() Styles {
	return Styles{
		Item:         lipgloss.NewStyle().Border(tabBorder, true).Padding(0, 1),
		SelectedItem: lipgloss.NewStyle().Border(activeTabBorder, true).Padding(0, 1),
		Filler:       lipgloss.NewStyle().Border(tabBorder, false, false, true, false).Padding(0, 1),
	}
}

type Styles struct {
	Item         lipgloss.Style
	SelectedItem lipgloss.Style
	Filler       lipgloss.Style
}

var (
	activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}
)
