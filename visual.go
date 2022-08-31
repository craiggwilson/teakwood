package teakwood

import tea "github.com/charmbracelet/bubbletea"

type Visual interface {
	tea.Model

	ViewWithBounds(Rectangle) string
}
