package specification

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
	value Satisfier[T]
}

func NewCompositeSpecification[T any](satisfier Satisfier[T]) CompositeSpecification[T] {
	return CompositeSpecification[T]{value: satisfier}
}
func (c CompositeSpecification[T]) IsSatisfiedBy(value T) bool {
	return c.value.IsSatisfiedBy(value)
}

func (c CompositeSpecification[T]) And(other Specification[T]) Specification[T] {
	return NewAndSpecification[T](c, other)
}

func (c CompositeSpecification[T]) AndNot(other Specification[T]) Specification[T] {
	return NewAndNotSpecification[T](c, other)
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

// AndSpecification
type AndSpecficiation[T any] struct {
	CompositeSpecification[T]
	leftCondition  Specification[T]
	rightCondition Specification[T]
}

func (and AndSpecficiation[T]) IsSatisfiedBy(value T) bool {
	return and.leftCondition.IsSatisfiedBy(value) && and.rightCondition.IsSatisfiedBy(value)
}
func NewAndSpecification[T any](left, right Specification[T]) AndSpecficiation[T] {
	return AndSpecficiation[T]{leftCondition: left, rightCondition: right}
}

type AndNotSpecficiation[T any] struct {
	CompositeSpecification[T]
	leftCondition  Specification[T]
	rightCondition Specification[T]
}

func (andn AndNotSpecficiation[T]) IsSatisfiedBy(value T) bool {
	return !(andn.leftCondition.IsSatisfiedBy(value) && andn.rightCondition.IsSatisfiedBy(value))
}
func NewAndNotSpecification[T any](left, right Specification[T]) AndNotSpecficiation[T] {
	return AndNotSpecficiation[T]{leftCondition: left, rightCondition: right}
}

type NotSpecficiation[T any] struct {
	CompositeSpecification[T]
	condition Specification[T]
}

func (ns NotSpecficiation[T]) IsSatisfiedBy(value T) bool {
	return !(ns.condition.IsSatisfiedBy(value))
}
func NewNotSpecification[T any](condition Specification[T]) NotSpecficiation[T] {
	return NotSpecficiation[T]{condition: condition, CompositeSpecification: CompositeSpecification[T]{value: condition}}
}

type OrNotSpecficiation[T any] struct {
	CompositeSpecification[T]
	leftCondition  Specification[T]
	rightCondition Specification[T]
}

func (on OrNotSpecficiation[T]) IsSatisfiedBy(value T) bool {
	return !(on.leftCondition.IsSatisfiedBy(value) || on.rightCondition.IsSatisfiedBy(value))
}
func NewOrNotSpecification[T any](left, right Specification[T]) OrNotSpecficiation[T] {
	return OrNotSpecficiation[T]{leftCondition: left, rightCondition: right}
}

type OrSpecficiation[T any] struct {
	CompositeSpecification[T]
	leftCondition  Specification[T]
	rightCondition Specification[T]
}

func (o OrSpecficiation[T]) IsSatisfiedBy(value T) bool {
	return o.leftCondition.IsSatisfiedBy(value) || o.rightCondition.IsSatisfiedBy(value)
}
func NewOrSpecification[T any](left, right Specification[T]) OrSpecficiation[T] {
	return OrSpecficiation[T]{leftCondition: left, rightCondition: right}
}
