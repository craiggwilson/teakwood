package iter

func Empty[S any]() Iterer[S] {
	return Generate(func() (S, bool) {
		var def S
		return def, false
	})
}
