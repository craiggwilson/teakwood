package tabs

type Opt func(*Model)

func WithStyles(styles Styles) Opt {
	return func(m *Model) {
		m.styles = styles
	}
}
