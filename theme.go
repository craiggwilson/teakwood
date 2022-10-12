package teakwood

import "github.com/charmbracelet/lipgloss"

type Styler interface {
	Style(string) (lipgloss.Style, bool)
}
