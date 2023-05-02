package specp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	ID        uint
	Username  string
	Firstname string
	Lastname  string
	Age       uint
	IsMale    bool
}

type OlderThanSpecification[T User] struct {
	CompositeSpecification[User]
	age uint
}

func (o OlderThanSpecification[T]) IsSatisfiedBy(u User) bool {
	return u.Age > o.age
}

type YoungerThanSpecification[T User] struct {
	CompositeSpecification[User]
	age uint
}

func (o YoungerThanSpecification[T]) IsSatisfiedBy(u User) bool {
	return u.Age < o.age
}

func NewYoungerThanSpecification(age uint) YoungerThanSpecification[User] {
	spec := YoungerThanSpecification[User]{age: age}
	spec.Value = NewCompositeSpecification[User](spec)
	return spec
}

func NewOlderThanSpecification(age uint) OlderThanSpecification[User] {
	spec := OlderThanSpecification[User]{age: age}
	spec.Value = NewCompositeSpecification[User](spec)
	return spec
}

type ExpectSexSpecification[T User] struct {
	CompositeSpecification[User]
	expectMale bool
}

func (o ExpectSexSpecification[T]) IsSatisfiedBy(u User) bool {
	return u.IsMale == o.expectMale
}

func NewExpectSexSpecification(isMale bool) ExpectSexSpecification[User] {
	spec := ExpectSexSpecification[User]{expectMale: isMale}
	spec.Value = NewCompositeSpecification[User](spec)
	return spec
}

func TestBasicRun(t *testing.T) {
	user := User{Age: 30, IsMale: true}
	olderSpec := NewOlderThanSpecification(20)
	youngerSpec := NewYoungerThanSpecification(40)
	alotOlderSpec := NewOlderThanSpecification(50)

	rs := olderSpec.IsSatisfiedBy(user)
	assert.True(t, rs)

	rs = youngerSpec.IsSatisfiedBy(user)
	assert.True(t, rs)

	combined := olderSpec.And(youngerSpec)
	rs = combined.IsSatisfiedBy(user)
	assert.True(t, rs)

	moreCombined := combined.And(alotOlderSpec)
	rs = moreCombined.IsSatisfiedBy(user)
	assert.False(t, rs)
}

func TestAndSpecification(t *testing.T) {
	user := User{Age: 20, IsMale: true}

	type test = struct {
		user           User
		name           string
		specfication   Specification[User]
		expectedResult bool
	}

	tests := []test{
		{
			user:           user,
			name:           "simple true answer",
			expectedResult: true,
			specfication:   NewOlderThanSpecification(18).And(NewOlderThanSpecification(10)),
		},
		{
			user:           user,
			name:           "simple false answer",
			expectedResult: false,
			specfication:   NewOlderThanSpecification(18).And(NewOlderThanSpecification(30)),
		},
		{
			user:           user,
			name:           "simple true multiple answer",
			expectedResult: true,
			specfication:   NewOlderThanSpecification(18).And(NewExpectSexSpecification(true)),
		},
		{
			user:           user,
			name:           "simple false multiple answer",
			expectedResult: false,
			specfication:   NewOlderThanSpecification(18).And(NewExpectSexSpecification(false)),
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			result := v.specfication.IsSatisfiedBy(user)
			assert.Equal(t, v.expectedResult, result)
		})
	}
}

func TestAndNotSpecification(t *testing.T) {
	user := User{Age: 20, IsMale: true}

	type test = struct {
		user           User
		name           string
		specfication   Specification[User]
		expectedResult bool
	}

	tests := []test{
		{
			user:           user,
			name:           "simple false answer",
			expectedResult: false,
			specfication:   NewOlderThanSpecification(18).AndNot(NewOlderThanSpecification(10)),
		},
		{
			user:           user,
			name:           "simple true answer",
			expectedResult: true,
			specfication:   NewOlderThanSpecification(18).AndNot(NewOlderThanSpecification(30)),
		},
		{
			user:           user,
			name:           "simple false multiple answer",
			expectedResult: false,
			specfication:   NewOlderThanSpecification(18).AndNot(NewExpectSexSpecification(true)),
		},
		{
			user:           user,
			name:           "simple true multiple answer",
			expectedResult: true,
			specfication:   NewOlderThanSpecification(18).AndNot(NewExpectSexSpecification(false)),
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			result := v.specfication.IsSatisfiedBy(user)
			assert.Equal(t, v.expectedResult, result)
		})
	}
}

func TestOrSpecification(t *testing.T) {
	user := User{Age: 20, IsMale: true}

	type test = struct {
		user           User
		name           string
		specfication   Specification[User]
		expectedResult bool
	}

	tests := []test{
		{
			user:           user,
			name:           "simple true answer",
			expectedResult: true,
			specfication:   NewOlderThanSpecification(25).Or(NewOlderThanSpecification(10)),
		},
		{
			user:           user,
			name:           "simple false answer",
			expectedResult: false,
			specfication:   NewOlderThanSpecification(25).Or(NewOlderThanSpecification(30)),
		},
		{
			user:           user,
			name:           "simple true multiple answer",
			expectedResult: true,
			specfication:   NewOlderThanSpecification(35).Or(NewExpectSexSpecification(true)),
		},
		{
			user:           user,
			name:           "simple false multiple answer",
			expectedResult: false,
			specfication:   NewOlderThanSpecification(30).Or(NewExpectSexSpecification(false)),
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			result := v.specfication.IsSatisfiedBy(user)
			assert.Equal(t, v.expectedResult, result)
		})
	}
}

func TestOrNotSpecification(t *testing.T) {
	user := User{Age: 20, IsMale: true}

	type test = struct {
		user           User
		name           string
		specfication   Specification[User]
		expectedResult bool
	}

	tests := []test{
		{
			user:           user,
			name:           "simple true answer",
			expectedResult: true,
			specfication:   NewOlderThanSpecification(25).OrNot(NewYoungerThanSpecification(18)),
		},
		{
			user:           user,
			name:           "simple false answer",
			expectedResult: false,
			specfication:   NewOlderThanSpecification(30).OrNot(NewYoungerThanSpecification(25)),
		},
		{
			user:           user,
			name:           "simple true multiple answer",
			expectedResult: true,
			specfication:   NewOlderThanSpecification(35).OrNot(NewExpectSexSpecification(false)),
		},
		{
			user:           user,
			name:           "simple false multiple answer",
			expectedResult: false,
			specfication:   NewOlderThanSpecification(35).OrNot(NewExpectSexSpecification(true)),
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			result := v.specfication.IsSatisfiedBy(user)
			assert.Equal(t, v.expectedResult, result)
		})
	}
}

func TestNotSpecification(t *testing.T) {
	user := User{Age: 20, IsMale: true}

	type test = struct {
		user           User
		name           string
		specfication   Specification[User]
		expectedResult bool
	}

	tests := []test{
		{
			user:           user,
			name:           "simple true answer",
			expectedResult: true,
			specfication:   NewOlderThanSpecification(25).Not(),
		},
		{
			user:           user,
			name:           "simple false answer",
			expectedResult: false,
			specfication:   NewOlderThanSpecification(18).Not(),
		},
		{
			user:           user,
			name:           "simple true multiple answer",
			expectedResult: true,
			specfication:   NewExpectSexSpecification(true).And(NewOlderThanSpecification(35).Not()),
		},
		{
			user:           user,
			name:           "simple false multiple answer",
			expectedResult: false,
			specfication:   NewExpectSexSpecification(true).And(NewOlderThanSpecification(18).Not()),
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			result := v.specfication.IsSatisfiedBy(user)
			assert.Equal(t, v.expectedResult, result)
		})
	}
}
