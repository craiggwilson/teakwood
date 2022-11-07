package label

type Opt func(*Model)

func WithClasses(classes ...string) Opt {
	return func(m *Model) {
		m.Visual.AddClasses(classes...)
	}
}

func WithHidden() Opt {
	return func(m *Model) {
		m.Visual.SetVisible(false)
	}
}

func WithID(id string) Opt {
	return func(m *Model) {
		m.Visual.SetID(id)
	}
}
