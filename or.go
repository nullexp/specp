package specp

// OrSpecification
type OrSpecificiation[T any] struct {
	CompositeSpecification[T]
	leftCondition  Specification[T]
	rightCondition Specification[T]
}

func NewOrSpecification[T any](left, right Specification[T]) OrSpecificiation[T] {
	return OrSpecificiation[T]{leftCondition: left, rightCondition: right}
}

func (Or OrSpecificiation[T]) IsSatisfiedBy(value T) bool {
	return Or.leftCondition.IsSatisfiedBy(value) || Or.rightCondition.IsSatisfiedBy(value)
}

func (c OrSpecificiation[T]) OrNot(other Specification[T]) Specification[T] {
	return NewOrNotSpecification[T](c, other)
}

func (c OrSpecificiation[T]) Not() Specification[T] {
	return NewNotSpecification[T](c)
}

func (c OrSpecificiation[T]) Or(other Specification[T]) Specification[T] {
	return NewOrSpecification[T](c, other)
}

func (c OrSpecificiation[T]) And(other Specification[T]) Specification[T] {
	return NewAndSpecificiation[T](c, other)
}
