package domain

// Automation is an interface for automations.
type Automation interface {
	SetInProgress() error
	SetInPending() error
	SetComplete() error
}

// IssueAutomation is an interface for automations that are triggered when an issue is assigned to me.
type IssueAutomation interface {
	Automation
}

// PrAssignedToMeAutomation is an interface for automations that are triggered when a PR is assigned to me.
type PrAssignedToMeAutomation interface {
	Automation
}

// PrAssignedToReviewerAutomation is an interface for automations that are triggered when a PR is assigned to a reviewer.
type PrAssignedToReviewerAutomation interface {
	Automation
}
