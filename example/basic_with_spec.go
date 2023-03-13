package example

// AccpetedForIntreview Without Specification
func AccpetedForIntreview(application ApplicationRequest) (accepted bool) {

	return OlderThan(18).
		And(YoungerThan(30)).
		And(NameNotEmpty(application.Firstname, application.Lastname)).
		And(ApplyTimeOlderThan(Week)).
		And(AllSkillExist(Golang, Docker)).
		And(HasMoreRelevantExperienceThan(4)).IsSatisfiedBy(application)

}
