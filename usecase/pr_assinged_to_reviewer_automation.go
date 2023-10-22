package usecase

import "github.com/igsr5/github-project-automation/domain"

type reviewPrAutomation struct{}

// NewReviewPrAutomation is a factory method to create a new instance of ReviewPrAutomation.
func NewReviewPrAutomation() domain.ReviewPrAutomation {
	return &reviewPrAutomation{}
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
