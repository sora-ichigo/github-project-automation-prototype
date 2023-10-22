package usecase

import "github.com/igsr5/github-project-automation/domain"

type prAssignedToMeAutomationImpl struct {
}

// NewPrAssignedToMeAutomation is a factory method to create a new instance of PrAssignedToMeAutomation.
func NewPrAssignedToMeAutomation() domain.PrAssignedToMeAutomation {
	return &prAssignedToMeAutomationImpl{}
}

func (a *prAssignedToMeAutomationImpl) SetInProgress() error {
	return nil
}

func (a *prAssignedToMeAutomationImpl) SetInPending() error {
	return nil
}

func (a *prAssignedToMeAutomationImpl) SetComplete() error {
	return nil
}
