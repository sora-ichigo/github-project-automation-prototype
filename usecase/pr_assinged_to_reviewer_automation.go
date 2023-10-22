package usecase

import "github.com/igsr5/github-project-automation/domain"

type prAssignedToReviewerAutomationImpl struct{}

// NewPrAssignedToReviewerAutomation is a factory method to create a new instance of PrAssignedToReviewerAutomation.
func NewPrAssignedToReviewerAutomation() domain.PrAssignedToReviewerAutomation {
	return &prAssignedToReviewerAutomationImpl{}
}

func (a *prAssignedToReviewerAutomationImpl) SetInProgress() error {
	// TODO: Implement
	return nil
}

func (a *prAssignedToReviewerAutomationImpl) SetInPending() error {
	// TODO: Implement
	return nil
}

func (a *prAssignedToReviewerAutomationImpl) SetComplete() error {
	// TODO: Implement
	return nil
}
