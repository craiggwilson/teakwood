package label

type Opt func(*Model)

func WithHidden() Opt {
	return func(m *Model) {
		m.visible = false
	}
}

func WithName(name string) Opt {
	return func(m *Model) {
		m.name = name
	}
}
