package specp

// AndSpecification
type AndNotSpecificiation[T any] struct {
	CompositeSpecification[T]
	leftCondition  Specification[T]
	rightCondition Specification[T]
}

func NewAndNotSpecificiation[T any](left, right Specification[T]) AndNotSpecificiation[T] {
	return AndNotSpecificiation[T]{leftCondition: left, rightCondition: right}
}

func (and AndNotSpecificiation[T]) IsSatisfiedBy(value T) bool {
	return !(and.leftCondition.IsSatisfiedBy(value) && and.rightCondition.IsSatisfiedBy(value))
}

func (c AndNotSpecificiation[T]) AndNot(other Specification[T]) Specification[T] {
	return NewAndNotSpecificiation[T](c, other)
}

func (c AndNotSpecificiation[T]) Not() Specification[T] {
	return NewNotSpecification[T](c)
}

func (c AndNotSpecificiation[T]) OrNot(other Specification[T]) Specification[T] {
	return NewOrNotSpecification[T](c, other)
}

func (c AndNotSpecificiation[T]) Or(other Specification[T]) Specification[T] {
	return NewOrSpecification[T](c, other)
}

func (c AndNotSpecificiation[T]) And(other Specification[T]) Specification[T] {
	return NewAndNotSpecificiation[T](c, other)
}
