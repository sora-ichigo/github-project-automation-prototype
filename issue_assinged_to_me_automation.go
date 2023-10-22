package main

type issueAssignedToMeAutomationImpl struct {
	issuefetcher    IssueFetcher
	projectV2Setter ProjectV2Setter
}

// NewIssueAssignedToMeAutomation is a factory method to create a new instance of IssueAssignedToMeAutomation.
func NewIssueAssignedToMeAutomation(issueFetcher IssueFetcher, projectV2Setter ProjectV2Setter) IssueAssignedToMeAutomation {
	return &issueAssignedToMeAutomationImpl{
		issuefetcher:    issueFetcher,
		projectV2Setter: projectV2Setter,
	}
}

func (a *issueAssignedToMeAutomationImpl) setInProgress() error {
	issues, err := a.issuefetcher.MyIssues()
	if err != nil {
		return err
	}

	categoryId := "1234"
	statusId := "5678"

	projectItems := make([]ProjectItem, 0, len(issues))
	for _, i := range issues {
		projectItems = append(projectItems, ProjectItem{Url: i.Url})
	}

	a.projectV2Setter.Set(categoryId, statusId, projectItems)
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
