package label

type Opt func(*Model)

func WithHidden() Opt {
	return func(m *Model) {
		m.widget.Visible = false
	}
}

func WithName(name string) Opt {
	return func(m *Model) {
		m.widget.Name = name
	}
}
