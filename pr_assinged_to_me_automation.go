package main

type prAssignedToMeAutomationImpl struct {
}

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
