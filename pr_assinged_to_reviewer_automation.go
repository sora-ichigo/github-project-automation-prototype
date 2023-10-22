package main

type prAssignedToReviewerAutomationImpl struct{}

// NewPrAssignedToReviewerAutomation is a factory method to create a new instance of PrAssignedToReviewerAutomation.
func NewPrAssignedToReviewerAutomation() PrAssignedToReviewerAutomation {
	return &prAssignedToReviewerAutomationImpl{}
}

func (a *prAssignedToReviewerAutomationImpl) setInProgress() error {
	return nil
}

func (a *prAssignedToReviewerAutomationImpl) setInPending() error {
	return nil
}

func (a *prAssignedToReviewerAutomationImpl) setComplete() error {
	return nil
}
