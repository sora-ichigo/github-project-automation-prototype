package main

type Automation interface {
	setInProgress() error
	setInPending() error
	setComplete() error
}

type IssueAssignedToMeAutomation interface {
	Automation
}

type PrAssignedToMeAutomation interface {
	Automation
}

type PrAssignedToReviewerAutomation interface {
	Automation
}
