package teakwood

import "github.com/charmbracelet/lipgloss"

type StyleSheet interface {
	Query(string) lipgloss.Style
}
