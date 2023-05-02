package specp

type NotSpecficiation[T any] struct {
	CompositeSpecification[T]
	condition Specification[T]
}

func NewNotSpecification[T any](condition Specification[T]) NotSpecficiation[T] {
	return NotSpecficiation[T]{condition: condition, CompositeSpecification: CompositeSpecification[T]{Value: condition}}
}

func (ns NotSpecficiation[T]) IsSatisfiedBy(value T) bool {
	return !(ns.condition.IsSatisfiedBy(value))
}

func (c NotSpecficiation[T]) OrNot(other Specification[T]) Specification[T] {
	return NewOrNotSpecification[T](c, other)
}

func (c NotSpecficiation[T]) Not() Specification[T] {
	return NewNotSpecification[T](c)
}

func (c NotSpecficiation[T]) Or(other Specification[T]) Specification[T] {
	return NewOrSpecification[T](c, other)
}

func (c NotSpecficiation[T]) AndNot(other Specification[T]) Specification[T] {
	return NewAndNotSpecificiation[T](c, other)
}

func (c NotSpecficiation[T]) And(other Specification[T]) Specification[T] {
	return NewAndSpecificiation[T](c, other)
}
