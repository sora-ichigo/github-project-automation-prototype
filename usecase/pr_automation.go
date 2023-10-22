package usecase

import "github.com/igsr5/github-project-automation/domain"

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
