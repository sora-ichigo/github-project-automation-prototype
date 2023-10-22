package command

import (
	"errors"

	usecase "github.com/igsr5/github-project-automation/usecase"
)

type projectV2SetterImpl struct {
}

// NewProjectV2Setter is a factory method to create a new instance of ProjectV2Setter.
func NewProjectV2Setter() usecase.ProjectV2Setter {
	return &projectV2SetterImpl{}
}

// Set sets project items.
func (s *projectV2SetterImpl) Set(categoryID string, statusID string, projectItems []usecase.ProjectItem) error {
	// 1. Add project items
	ids := make([]projectID, 0, len(projectItems))
	for _, item := range projectItems {
		id, err := addProjectItem(item.URL)
		if err != nil {
			return errors.Join(err)
		}
		ids = append(ids, id)
	}

	// 2. Move project item status to specified status
	for _, id := range ids {
		err := moveStatusProjectItem(id, statusID)
		if err != nil {
			return errors.Join(err)
		}
	}

	// 3. Move project item category to specified category
	for _, id := range ids {
		err := moveCategoryProjectItem(id, categoryID)
		if err != nil {
			return errors.Join(err)
		}
	}

	return nil
}

type projectID string

func addProjectItem(url string) (projectID, error) {
	// TODO: implement
	return "", nil
}

func moveStatusProjectItem(id projectID, statusID string) error {
	// TODO: implement
	return nil
}

func moveCategoryProjectItem(id projectID, categoryID string) error {
	// TODO: implement
	return nil
}
