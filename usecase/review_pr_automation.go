package usecase

import (
	"fmt"
	"log"

	variables "github.com/sora-ichigo/github-project-automation"
	"github.com/sora-ichigo/github-project-automation/domain"
)

type reviewPrAutomation struct {
	reviewPrFetcher ReviewPrFetcher
	projectV2Setter ProjectV2Setter
}

// NewReviewPrAutomation is a factory method to create a new instance of ReviewPrAutomation.
func NewReviewPrAutomation(reviewPrFetcher ReviewPrFetcher, projectV2Setter ProjectV2Setter) domain.ReviewPrAutomation {
	return &reviewPrAutomation{
		reviewPrFetcher: reviewPrFetcher,
		projectV2Setter: projectV2Setter,
	}
}

func (a *reviewPrAutomation) SetInProgress() error {
	prs, err := a.reviewPrFetcher.UnReviewedPrs()
	if err != nil {
		return fmt.Errorf("ReviewPRAutomation SetInProgress is failed: %w", err)
	}

	categoryID := variables.REVIEW_PR_CATEGORY_ID
	statusID := variables.IN_PROGRESS_STATUS_ID

	projectItems := make([]ProjectItem, 0, len(prs))
	for _, i := range prs {
		projectItems = append(projectItems, ProjectItem{URL: i.URL})
	}

	log.Printf("review pr automation\n")
	log.Printf("Found %d target review prs in SetInProgress.\n", len(projectItems))

	err = a.projectV2Setter.Set(categoryID, statusID, projectItems)
	if err != nil {
		return fmt.Errorf("ReviewPRAutomation SetInProgress is failed: %w", err)
	}

	return nil
}

func (a *reviewPrAutomation) SetInPending() error {
	prs, err := a.reviewPrFetcher.ReviewedPrs()
	if err != nil {
		return fmt.Errorf("ReviewPRAutomation SetInPending is failed: %w", err)
	}

	categoryID := variables.REVIEW_PR_CATEGORY_ID
	statusID := variables.IN_PENDING_STATUS_ID

	projectItems := make([]ProjectItem, 0, len(prs))
	for _, i := range prs {
		projectItems = append(projectItems, ProjectItem{URL: i.URL})
	}

	log.Printf("review pr automation\n")
	log.Printf("Found %d target review prs in SetInPending.\n", len(projectItems))

	err = a.projectV2Setter.Set(categoryID, statusID, projectItems)
	if err != nil {
		return fmt.Errorf("ReviewPRAutomation SetInPending is failed: %w", err)
	}

	return nil
}

func (a *reviewPrAutomation) SetComplete() error {
	prs, err := a.reviewPrFetcher.ApprovedPrs()
	if err != nil {
		return fmt.Errorf("ReviewPRAutomation SetInComplete is failed: %w", err)
	}

	categoryID := variables.REVIEW_PR_CATEGORY_ID
	statusID := variables.COMPLETE_STATUS_ID

	projectItems := make([]ProjectItem, 0, len(prs))
	for _, i := range prs {
		projectItems = append(projectItems, ProjectItem{URL: i.URL})
	}

	log.Printf("review pr automation\n")
	log.Printf("Found %d target review prs in SetComplete.\n", len(projectItems))

	err = a.projectV2Setter.Set(categoryID, statusID, projectItems)
	if err != nil {
		return fmt.Errorf("ReviewPRAutomation SetInComplete is failed: %w", err)
	}

	return nil
}
