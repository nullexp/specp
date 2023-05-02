package specp

type Satisfier[T any] interface {
	IsSatisfiedBy(value T) bool
}

type Specification[T any] interface {
	Satisfier[T]
	And(other Specification[T]) Specification[T]
	AndNot(other Specification[T]) Specification[T]
	Not() Specification[T]
	OrNot(other Specification[T]) Specification[T]
	Or(other Specification[T]) Specification[T]
}

type CompositeSpecification[T any] struct {
	Value Satisfier[T]
}

func NewCompositeSpecification[T any](satisfier Satisfier[T]) CompositeSpecification[T] {
	return CompositeSpecification[T]{Value: satisfier}
}

func (c CompositeSpecification[T]) IsSatisfiedBy(value T) bool {
	return c.Value.IsSatisfiedBy(value)
}

func (c CompositeSpecification[T]) And(other Specification[T]) Specification[T] {
	return NewAndSpecificiation[T](c, other)
}

func (c CompositeSpecification[T]) AndNot(other Specification[T]) Specification[T] {
	return NewAndNotSpecificiation[T](c, other)
}

func (c CompositeSpecification[T]) Not() Specification[T] {
	return NewNotSpecification[T](c)
}

func (c CompositeSpecification[T]) OrNot(other Specification[T]) Specification[T] {
	return NewOrNotSpecification[T](c, other)
}

func (c CompositeSpecification[T]) Or(other Specification[T]) Specification[T] {
	return NewOrSpecification[T](c, other)
}
