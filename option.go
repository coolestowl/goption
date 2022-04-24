package goption

type Option[T any] interface {
	Apply(*T)
}

func FuncOption[T any](f func(*T)) Option[T] {
	return fnOption[T](f)
}

type fnOption[T any] func(*T)

func (f fnOption[T]) Apply(t *T) {
	f(t)
}
