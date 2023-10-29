package usecase

import (
	"fmt"
	"log"

	variables "github.com/igsr5/github-project-automation"
	"github.com/igsr5/github-project-automation/domain"
)

type prAutomationImpl struct {
	prFetcher       PrFetcher
	projectV2Setter ProjectV2Setter
}

// NewPrAutomation is a factory method to create a new instance of PrAutomation.
func NewPrAutomation(prFetcher PrFetcher, projectV2Setter ProjectV2Setter) domain.PrAutomation {
	return &prAutomationImpl{
		prFetcher:       prFetcher,
		projectV2Setter: projectV2Setter,
	}
}

// SetInProgress sets pull requests that are in progress.
func (a *prAutomationImpl) SetInProgress() error {
	prs, err := a.prFetcher.UnReviewedPrs()
	if err != nil {
		return fmt.Errorf("PRAutomation SetInProgress is failed: %w", err)
	}

	categoryID := variables.PR_CATEGORY_ID
	statusID := variables.IN_PROGRESS_STATUS_ID

	projectItems := make([]ProjectItem, 0, len(prs))
	for _, i := range prs {
		projectItems = append(projectItems, ProjectItem{URL: i.URL})
	}

	log.Printf("pr automation\n")
	log.Printf("Found %d target prs in SetInProgress.\n", len(projectItems))
	a.projectV2Setter.Set(categoryID, statusID, projectItems)
	return nil
}

// SetInPending sets pull requests that are requested to review.
func (a *prAutomationImpl) SetInPending() error {
	prs, err := a.prFetcher.ReviewedPrs()
	if err != nil {
		return fmt.Errorf("PRAutomation SetInPending is failed: %w", err)
	}

	categoryID := variables.PR_CATEGORY_ID
	statusID := variables.IN_PENDING_STATUS_ID

	projectItems := make([]ProjectItem, 0, len(prs))
	for _, i := range prs {
		projectItems = append(projectItems, ProjectItem{URL: i.URL})
	}

	log.Printf("pr automation\n")
	log.Printf("Found %d target prs in SetInPending.\n", len(projectItems))
	a.projectV2Setter.Set(categoryID, statusID, projectItems)
	return nil
}

// SetComplete sets pull requests that are approved.
func (a *prAutomationImpl) SetComplete() error {
	prs, err := a.prFetcher.ApprovedPrs()
	if err != nil {
		return fmt.Errorf("PRAutomation SetComplete is failed: %w", err)
	}

	categoryID := variables.PR_CATEGORY_ID
	statusID := variables.COMPLETE_STATUS_ID

	projectItems := make([]ProjectItem, 0, len(prs))
	for _, i := range prs {
		projectItems = append(projectItems, ProjectItem{URL: i.URL})
	}

	log.Printf("pr automation\n")
	log.Printf("Found %d target prs in SetComplete.\n", len(projectItems))
	a.projectV2Setter.Set(categoryID, statusID, projectItems)
	return nil
}
