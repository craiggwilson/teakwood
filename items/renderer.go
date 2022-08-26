package items

type Renderer[T any] interface {
	Render(T) string
}

type RenderFunc[T any] func(T) string

func (r RenderFunc[T]) Render(item T) string {
	return r(item)
}
