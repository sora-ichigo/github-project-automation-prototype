package usecase

import (
	variables "github.com/igsr5/github-project-automation"
	"github.com/igsr5/github-project-automation/domain"
)

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

	categoryID := variables.ISSUE_CATEGORY_ID
	statusID := variables.IN_PROGRESS_STATUS_ID

	projectItems := make([]ProjectItem, 0, len(issues))
	for _, i := range issues {
		projectItems = append(projectItems, ProjectItem{URL: i.URL})
	}

	a.projectV2Setter.Set(categoryID, statusID, projectItems)
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
