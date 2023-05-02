package specp

// AndSpecification
type AndSpecificiation[T any] struct {
	CompositeSpecification[T]
	leftCondition  Specification[T]
	rightCondition Specification[T]
}

func NewAndSpecificiation[T any](left, right Specification[T]) AndSpecificiation[T] {
	return AndSpecificiation[T]{leftCondition: left, rightCondition: right}
}

func (and AndSpecificiation[T]) IsSatisfiedBy(value T) bool {
	return and.leftCondition.IsSatisfiedBy(value) && and.rightCondition.IsSatisfiedBy(value)
}

func (c AndSpecificiation[T]) AndNot(other Specification[T]) Specification[T] {
	return NewAndNotSpecificiation[T](c, other)
}

func (c AndSpecificiation[T]) Not() Specification[T] {
	return NewNotSpecification[T](c)
}

func (c AndSpecificiation[T]) OrNot(other Specification[T]) Specification[T] {
	return NewOrNotSpecification[T](c, other)
}

func (c AndSpecificiation[T]) Or(other Specification[T]) Specification[T] {
	return NewOrSpecification[T](c, other)
}

func (c AndSpecificiation[T]) And(other Specification[T]) Specification[T] {
	return NewAndSpecificiation[T](c, other)
}
