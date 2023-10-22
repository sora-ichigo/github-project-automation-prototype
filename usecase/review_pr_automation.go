package usecase

import "github.com/igsr5/github-project-automation/domain"

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
	// TODO: Implement
	return nil
}

func (a *reviewPrAutomation) SetInPending() error {
	// TODO: Implement
	return nil
}

func (a *reviewPrAutomation) SetComplete() error {
	// TODO: Implement
	return nil
}
