package main

type prAssignedToMeAutomationImpl struct {
}

// NewPrAssignedToMeAutomation is a factory method to create a new instance of PrAssignedToMeAutomation.
func NewPrAssignedToMeAutomation() PrAssignedToMeAutomation {
	return &prAssignedToMeAutomationImpl{}
}

func (a *prAssignedToMeAutomationImpl) setInProgress() error {
	return nil
}

func (a *prAssignedToMeAutomationImpl) setInPending() error {
	return nil
}

func (a *prAssignedToMeAutomationImpl) setComplete() error {
	return nil
}
