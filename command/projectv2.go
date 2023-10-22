package command

import (
	"encoding/json"
	"errors"
	"os/exec"

	variables "github.com/igsr5/github-project-automation"
	usecase "github.com/igsr5/github-project-automation/usecase"
)

type Item struct {
	ID string `json:"id"`
}

type projectV2SetterImpl struct {
}

// NewProjectV2Setter is a factory method to create a new instance of ProjectV2Setter.
func NewProjectV2Setter() usecase.ProjectV2Setter {
	return &projectV2SetterImpl{}
}

// Set sets project items.
func (s *projectV2SetterImpl) Set(categoryID string, statusID string, projectItems []usecase.ProjectItem) error {
	// 1. Add project items
	items := make([]Item, 0, len(projectItems))
	for _, item := range projectItems {
		item, err := addProjectItem(item.URL)
		if err != nil {
			return errors.Join(err)
		}
		items = append(items, *item)
	}

	// 2. Move project item status to specified status
	for _, id := range items {
		err := moveStatusProjectItem(id, statusID)
		if err != nil {
			return errors.Join(err)
		}
	}

	// 3. Move project item category to specified category
	for _, id := range items {
		err := moveCategoryProjectItem(id, categoryID)
		if err != nil {
			return errors.Join(err)
		}
	}

	return nil
}

func addProjectItem(url string) (*Item, error) {
	cmd := exec.Command("gh", "project", "item-add", variables.PROJECT_ID, "--owner", "wantedly", "--url", url, "--format", "json")
	output, err := cmd.Output()
	if err != nil {
		return nil, errors.Join(err)
	}

	var item Item
	if err := json.Unmarshal(output, &item); err != nil {
		return nil, errors.Join(err)
	}

	return &item, nil
}

func moveStatusProjectItem(item Item, statusID string) error {
	// TODO: implement
	return nil
}

func moveCategoryProjectItem(item Item, categoryID string) error {
	// TODO: implement
	return nil
}
