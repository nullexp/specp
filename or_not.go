package specp

// OrSpecification
type OrNotSpecificiation[T any] struct {
	CompositeSpecification[T]
	leftCondition  Specification[T]
	rightCondition Specification[T]
}

func NewOrNotSpecification[T any](left, right Specification[T]) OrNotSpecificiation[T] {
	return OrNotSpecificiation[T]{leftCondition: left, rightCondition: right}
}

func (or OrNotSpecificiation[T]) IsSatisfiedBy(value T) bool {
	return !(or.leftCondition.IsSatisfiedBy(value) || or.rightCondition.IsSatisfiedBy(value))
}
func (c OrNotSpecificiation[T]) OrNot(other Specification[T]) Specification[T] {
	return NewOrNotSpecification[T](c, other)
}

func (c OrNotSpecificiation[T]) Not() Specification[T] {
	return NewNotSpecification[T](c)
}

func (c OrNotSpecificiation[T]) Or(other Specification[T]) Specification[T] {
	return NewOrSpecification[T](c, other)
}

func (c OrNotSpecificiation[T]) AndNot(other Specification[T]) Specification[T] {
	return NewAndNotSpecificiation[T](c, other)
}

func (c OrNotSpecificiation[T]) And(other Specification[T]) Specification[T] {
	return NewAndSpecificiation[T](c, other)
}
