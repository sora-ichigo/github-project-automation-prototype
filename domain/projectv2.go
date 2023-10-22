package domain

// ProjectItem is a struct that represents a project item.
type ProjectItem struct {
	Url string
}

// ProjectV2Setter is an interface for setting project items.
type ProjectV2Setter interface {
	Set(categoryId string, statusId string, projectItems []ProjectItem) error
}
