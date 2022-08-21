package items

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/craiggwilson/teakwood/label"
)

type Source interface {
	Len() int
	Item(int) tea.Model
}

type ModelsSource[T tea.Model] []T

func (s ModelsSource[T]) Len() int {
	return len(s)
}

func (s ModelsSource[T]) Item(i int) tea.Model {
	return s[i]
}

type StringersSource[T fmt.Stringer] []T

func (s StringersSource[T]) Len() int {
	return len(s)
}

func (s StringersSource[T]) Item(i int) tea.Model {
	return label.New(s[i].String())
}

type StringsSource []string

func (s StringsSource) Len() int {
	return len(s)
}

func (s StringsSource) Item(i int) tea.Model {
	return label.New(s[i])
}
