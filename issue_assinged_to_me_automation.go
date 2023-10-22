package main

type issueAssignedToMeAutomationImpl struct {
	issuefetcher    *IssueFetcher
	projectV2Setter *ProjectV2Setter
}

// NewIssueAssignedToMeAutomation is a factory method to create a new instance of IssueAssignedToMeAutomation.
func NewIssueAssignedToMeAutomation(issueFetcher *IssueFetcher, projectV2Setter *ProjectV2Setter) IssueAssignedToMeAutomation {
	return &issueAssignedToMeAutomationImpl{
		issuefetcher:    issueFetcher,
		projectV2Setter: projectV2Setter,
	}
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
