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

// PrAutomation is an interface for automations that are triggered when a PR is assigned to me.
type PrAutomation interface {
	Automation
}

// ReviewPrAutomation is an interface for automations that are triggered when a PR is assigned to a reviewer.
type ReviewPrAutomation interface {
	Automation
}
