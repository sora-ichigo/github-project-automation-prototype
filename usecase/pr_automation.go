package usecase

import "github.com/igsr5/github-project-automation/domain"

type prAutomationImpl struct {
}

// NewPrAutomation is a factory method to create a new instance of PrAutomation.
func NewPrAutomation() domain.PrAutomation {
	return &prAutomationImpl{}
}

func (a *prAutomationImpl) SetInProgress() error {
	// TODO: Implement
	return nil
}

func (a *prAutomationImpl) SetInPending() error {
	// TODO: Implement
	return nil
}

func (a *prAutomationImpl) SetComplete() error {
	// TODO: Implement
	return nil
}
