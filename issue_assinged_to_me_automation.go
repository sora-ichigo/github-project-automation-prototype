package main

type issueAssignedToMeAutomationImpl struct{}

func NewIssueAssignedToMeAutomation() IssueAssignedToMeAutomation {
	return &issueAssignedToMeAutomationImpl{}
}

func (a *issueAssignedToMeAutomationImpl) setInProgress() error {
	return nil
}

// this status is not used, so do nothing
func (a *issueAssignedToMeAutomationImpl) setInPending() error {
	return nil
}

// this status is not used, so do nothing
func (a *issueAssignedToMeAutomationImpl) setComplete() error {
	return nil
}
