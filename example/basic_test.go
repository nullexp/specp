package example

import (
	"testing"
	"time"
)

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
var getTestCases = []struct {
	name         string
	application  func() ApplicationRequest
	wantAccepted bool
}{
	{
		name: "Shoud valid sample pass",
		application: func() (out ApplicationRequest) {
			out = validSample
			return
		},
		wantAccepted: true,
	},
	{
		name: "Shoud age be greater than 18",
		application: func() (out ApplicationRequest) {
			out = validSample
			out.Age = 16
			return
		},
		wantAccepted: false,
	},
	{
		name: "Shoud age be greater than 40",
		application: func() (out ApplicationRequest) {
			out = validSample
			out.Age = 40
			return
		},
		wantAccepted: false,
	},
	{
		name: "Shoud firstname not empty",
		application: func() (out ApplicationRequest) {
			out = validSample
			out.Firstname = ""
			return
		},
		wantAccepted: false,
	},
	{
		name: "Shoud last name not empty",
		application: func() (out ApplicationRequest) {
			out = validSample
			out.Lastname = ""
			return
		},
		wantAccepted: false,
	},
	{
		name: "Shoud apply time not be more than a week",
		application: func() (out ApplicationRequest) {
			out = validSample
			out.ApplyTime = now.Add(Week * 2)
			return
		},
		wantAccepted: false,
	},
	{
		name: "Shoud have required skills",
		application: func() (out ApplicationRequest) {
			out = validSample
			out.Skills = []Skill{}
			return
		},
		wantAccepted: false,
	},
	{
		name: "Shoud have atleast 4 years relevant experience",
		application: func() (out ApplicationRequest) {
			out = validSample
			out.RelevantExperience = 3
			return
		},
		wantAccepted: false,
	},
}

func TestIsAccpeted(t *testing.T) {

	for _, tt := range getTestCases {
		t.Run(tt.name, func(t *testing.T) {
			if gotAccepted := IsAccpeted(tt.application()); gotAccepted != tt.wantAccepted {
				t.Errorf("AccpetedForIntreview() = %v, want %v", gotAccepted, tt.wantAccepted)
			}
		})
	}
}

func TestIsAccpetedWithSpec(t *testing.T) {

	for _, tt := range getTestCases {
		t.Run(tt.name, func(t *testing.T) {
			if gotAccepted := IsAccpetedWithSpec(tt.application()); gotAccepted != tt.wantAccepted {
				t.Errorf("AccpetedForIntreview() = %v, want %v", gotAccepted, tt.wantAccepted)
			}
		})
	}
}
