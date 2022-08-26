package tabs

type Opt[T any] func(*Model[T])

func WithStyles[T any](styles Styles) Opt[T] {
	return func(m *Model[T]) {
		m.styles = styles
	}
}
