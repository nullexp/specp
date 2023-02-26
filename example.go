package specp

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Skill string

const ()

type ApplicationRequest struct {
	ID                 uuid.UUID
	Firstname          string
	Lastname           string
	ApplyTime          time.Time
	Age                int
	RelevantExperience int
}

const week = 7 * 24 * time.Hour

func isEmpty(value string) bool {
	return strings.Trim(value, " ") == ""
}

// AccpetedForIntreview Without Specification
// func AccpetedForIntreview(application ApplicationRequest) (accepted bool) {

// 	// Applicant age must be more and equal than 18 and less than 30
// 	if application.Age < 18 || application.Age > 30 {
// 		return
// 	}

// 	// Full name must be
// 	if isEmpty(application.Firstname) || isEmpty(application.Lastname) {
// 		return
// 	}

// 	// apply time must not be older than a week day
// 	if time.Since(application.ApplyTime) > week {
// 		return
// 	}

// }
