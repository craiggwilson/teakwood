package adapter

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/craiggwilson/teakwood/label"
)

func Strings(strings ...string) []tea.Model {
	models := make([]tea.Model, len(strings))
	for i := 0; i < len(strings); i++ {
		models[i] = label.New(strings[i])
	}
	return models
}
