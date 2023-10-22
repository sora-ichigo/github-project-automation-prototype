package usecase

import (
	variables "github.com/igsr5/github-project-automation"
	"github.com/igsr5/github-project-automation/domain"
)

type issueAutomationImpl struct {
	issuefetcher    IssueFetcher
	projectV2Setter ProjectV2Setter
}

// NewIssueAutomation is a factory method to create a new instance of IssueAutomation.
func NewIssueAutomation(issueFetcher IssueFetcher, projectV2Setter ProjectV2Setter) domain.IssueAutomation {
	return &issueAutomationImpl{
		issuefetcher:    issueFetcher,
		projectV2Setter: projectV2Setter,
	}
}

func (a *issueAutomationImpl) SetInProgress() error {
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
func (a *issueAutomationImpl) SetInPending() error {
	return nil
}

// this status is not used, so do nothing
func (a *issueAutomationImpl) SetComplete() error {
	return nil
}
