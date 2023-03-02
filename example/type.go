package example

import (
	"strings"
	"time"
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
)

const (
	Golang Skill = "Golang"
	Csharp Skill = "CSharp"
	Java   Skill = "Java"
	Rust   Skill = "Rust"
	Docker Skill = "Docker"
	Git    Skill = "Git"
	Week         = 7 * 24 * time.Hour
)

func isEmpty(value string) bool {
	return strings.Trim(value, " ") == ""
}

var (
	now                            = time.Now()
	validSample ApplicationRequest = ApplicationRequest{
		Firstname:          "Sam",
		Lastname:           "Smith",
		ApplyTime:          now,
		Age:                18,
		RelevantExperience: 4,
		Skills:             []Skill{Docker, Golang},
	}
)
