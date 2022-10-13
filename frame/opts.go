package frame

type Opt func(*Model)

func WithName(name string) Opt {
	return func(m *Model) {
		m.widget.Name = name
	}
}
