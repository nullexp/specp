package example

import (
	"time"
)

// IsAccpeted (for interview) Without Specification
func IsAccpeted(application ApplicationRequest) (accepted bool) {

	// Applicant age must be more and equal than 18 and less than 30
	if application.Age < 18 || application.Age > 30 {
		return
	}

	// Full name must not be empty
	if isEmpty(application.Firstname) || isEmpty(application.Lastname) {
		return
	}

	// apply time must not be older than a week day
	if time.Until(application.ApplyTime) > Week {
		return
	}

	// Creating a map for fast skill search
	skillExist := map[Skill]bool{}
	for _, v := range application.Skills {
		skillExist[v] = true
	}

	// We expect the applicant know Golang and Docker
	if !skillExist[Golang] || !skillExist[Docker] {
		return
	}

	// At least 4 years of relevant experience is expected
	if application.RelevantExperience < 4 {
		return
	}

	return true
}
