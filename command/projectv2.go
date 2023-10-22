package command

import (
	usecase "github.com/igsr5/github-project-automation/usecase"
)

type projectV2SetterImpl struct {
}

// NewProjectV2Setter is a factory method to create a new instance of ProjectV2Setter.
func NewProjectV2Setter() usecase.ProjectV2Setter {
	return &projectV2SetterImpl{}
}

// Set sets project items.
func (s *projectV2SetterImpl) Set(categoryId string, statusId string, projectItems []usecase.ProjectItem) error {
	// TODO: Implement
	return nil
}
