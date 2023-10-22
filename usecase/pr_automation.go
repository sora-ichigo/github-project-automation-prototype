package usecase

import (
	variables "github.com/igsr5/github-project-automation"
	"github.com/igsr5/github-project-automation/domain"
)

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
	prs, err := a.prFetcher.WorkInProgressPrs()
	if err != nil {
		return err
	}

	categoryID := variables.PR_CATEGORY_ID
	statusID := variables.IN_PROGRESS_STATUS_ID

	projectItems := make([]ProjectItem, 0, len(prs))
	for _, i := range prs {
		projectItems = append(projectItems, ProjectItem{URL: i.URL})
	}

	a.projectV2Setter.Set(categoryID, statusID, projectItems)
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
