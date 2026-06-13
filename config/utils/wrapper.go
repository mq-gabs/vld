package utils

type V[T any] struct {
	value *T
}

func W[T any](v T) V[T] {
	return V[T]{
		value: &v,
	}
}

func (p V[T]) IsSet() bool {
	return p.value != nil
}

func (p V[T]) Get() T {
	var zero T
	if p.value == nil {
		return zero
	}

	return *p.value
}
