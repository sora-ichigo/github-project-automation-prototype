package usecase

import "github.com/igsr5/github-project-automation/domain"

type issueAssignedToMeAutomationImpl struct {
	issuefetcher    IssueFetcher
	projectV2Setter ProjectV2Setter
}

// NewIssueAssignedToMeAutomation is a factory method to create a new instance of IssueAssignedToMeAutomation.
func NewIssueAssignedToMeAutomation(issueFetcher IssueFetcher, projectV2Setter ProjectV2Setter) domain.IssueAssignedToMeAutomation {
	return &issueAssignedToMeAutomationImpl{
		issuefetcher:    issueFetcher,
		projectV2Setter: projectV2Setter,
	}
}

func (a *issueAssignedToMeAutomationImpl) SetInProgress() error {
	issues, err := a.issuefetcher.MyIssues()
	if err != nil {
		return err
	}

	// TODO: Implement
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
func (a *issueAssignedToMeAutomationImpl) SetInPending() error {
	return nil
}

// this status is not used, so do nothing
func (a *issueAssignedToMeAutomationImpl) SetComplete() error {
	return nil
}
