package example

import (
	"strings"
	"time"

	"github.com/nullexp/specp"
)

type (
	Skill string

	ApplicationRequest struct {
		Firstname          string
		Lastname           string
		ApplyTime          time.Time
		Age                int
		RelevantExperience int
		Skills             []Skill
	}

	OlderThanSpec[T ApplicationRequest] struct {
		specp.CompositeSpecification[ApplicationRequest]
		age int
	}

	YoungerThanSpec[T ApplicationRequest] struct {
		specp.CompositeSpecification[ApplicationRequest]
		age int
	}

	NameNotEmptySpec[T ApplicationRequest] struct {
		specp.CompositeSpecification[ApplicationRequest]
		firstname, lastname string
	}

	ApplyTimeOlderThanSpec[T ApplicationRequest] struct {
		specp.CompositeSpecification[ApplicationRequest]
		duration time.Duration
	}

	AllSkillExistSpec[T ApplicationRequest] struct {
		specp.CompositeSpecification[ApplicationRequest]
		wantedSkills []Skill
	}

	HasMoreRelevantExperienceThanSpec[T ApplicationRequest] struct {
		specp.CompositeSpecification[ApplicationRequest]
		relevantExperience int
	}
)

const (
	Golang Skill = "Golang"
	Csharp Skill = "C#"
	Java   Skill = "Java"
	Rust   Skill = "Rust"
	Docker Skill = "Docker"
	Git    Skill = "Git"
	Week         = 7 * 24 * time.Hour
)

func isEmpty(value string) bool {
	return strings.Trim(value, " ") == ""
}

func (o OlderThanSpec[T]) IsSatisfiedBy(ar ApplicationRequest) bool {
	return ar.Age > o.age
}

func (y YoungerThanSpec[T]) IsSatisfiedBy(ar ApplicationRequest) bool {
	return ar.Age < y.age
}

func (n NameNotEmptySpec[T]) IsSatisfiedBy(ar ApplicationRequest) bool {
	return !isEmpty(ar.Firstname) && !isEmpty(ar.Lastname)
}

func (at ApplyTimeOlderThanSpec[T]) IsSatisfiedBy(ar ApplicationRequest) bool {
	return time.Until(ar.ApplyTime) < at.duration
}

func (as AllSkillExistSpec[T]) IsSatisfiedBy(ar ApplicationRequest) bool {
	skillExist := map[Skill]bool{}
	for _, v := range ar.Skills {
		skillExist[v] = true
	}

	for _, v := range as.wantedSkills {

		if !skillExist[v] {
			return false
		}
	}

	return true
}

func (hm HasMoreRelevantExperienceThanSpec[T]) IsSatisfiedBy(ar ApplicationRequest) bool {
	return hm.relevantExperience < ar.RelevantExperience
}

func OlderThan(age int) *OlderThanSpec[ApplicationRequest] {
	spec := OlderThanSpec[ApplicationRequest]{age: age}
	spec.Value = specp.NewCompositeSpecification[ApplicationRequest](spec)
	return &spec
}

func YoungerThan(age int) *YoungerThanSpec[ApplicationRequest] {
	spec := YoungerThanSpec[ApplicationRequest]{age: age}
	spec.Value = specp.NewCompositeSpecification[ApplicationRequest](spec)
	return &spec
}

func NameNotEmpty(firstname, lastname string) NameNotEmptySpec[ApplicationRequest] {
	spec := NameNotEmptySpec[ApplicationRequest]{firstname: firstname, lastname: lastname}
	spec.Value = specp.NewCompositeSpecification[ApplicationRequest](spec)
	return spec
}

func ApplyTimeOlderThan(duration time.Duration) ApplyTimeOlderThanSpec[ApplicationRequest] {
	spec := ApplyTimeOlderThanSpec[ApplicationRequest]{duration: duration}
	spec.Value = specp.NewCompositeSpecification[ApplicationRequest](spec)
	return spec
}

func AllSkillExist(wantedSkills ...Skill) AllSkillExistSpec[ApplicationRequest] {
	spec := AllSkillExistSpec[ApplicationRequest]{wantedSkills: wantedSkills}
	spec.Value = specp.NewCompositeSpecification[ApplicationRequest](spec)
	return spec
}

func HasMoreRelevantExperienceThan(relevantExperience int) HasMoreRelevantExperienceThanSpec[ApplicationRequest] {
	spec := HasMoreRelevantExperienceThanSpec[ApplicationRequest]{relevantExperience: relevantExperience}
	spec.Value = specp.NewCompositeSpecification[ApplicationRequest](spec)
	return spec
}
