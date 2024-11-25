package usecase

import (
	"fmt"
	"log"

	variables "github.com/sora-ichigo/github-project-automation"
	"github.com/sora-ichigo/github-project-automation/domain"
)

type issueAutomationImpl struct {
	issueFetcher    IssueFetcher
	projectV2Setter ProjectV2Setter
}

// NewIssueAutomation is a factory method to create a new instance of IssueAutomation.
func NewIssueAutomation(issueFetcher IssueFetcher, projectV2Setter ProjectV2Setter) domain.IssueAutomation {
	return &issueAutomationImpl{
		issueFetcher:    issueFetcher,
		projectV2Setter: projectV2Setter,
	}
}

func (a *issueAutomationImpl) SetInProgress() error {
	issues, err := a.issueFetcher.MyIssues()
	if err != nil {
		return fmt.Errorf("IssueAutomation SetInProgress is failed: %w", err)
	}

	categoryID := variables.ISSUE_CATEGORY_ID
	statusID := variables.IN_PROGRESS_STATUS_ID

	projectItems := make([]ProjectItem, 0, len(issues))
	for _, i := range issues {
		projectItems = append(projectItems, ProjectItem{URL: i.URL})
	}

	log.Printf("issues automation\n")
	log.Printf("Found %d target issues in SetInProgress.\n", len(projectItems))
	err = a.projectV2Setter.Set(categoryID, statusID, projectItems)
	if err != nil {
		return fmt.Errorf("IssueAutomation SetInProgress is failed: %w", err)
	}
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
