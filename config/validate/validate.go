package validate

type Validate[T any] func(T) error

type Validator[T any] interface {
	Validate(T) error
}
