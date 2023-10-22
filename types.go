package main

type Automation interface {
	setInProgress() error
	setInPending() error
	setComplete() error
}

type IssueAssignedToMeAutomation interface {
	Automation
}

type PRAssignedToMeAutomation interface {
	Automation
}

type PRAssignedToReviewerAutomation interface {
	Automation
}
