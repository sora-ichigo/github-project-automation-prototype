package usecase

// ProjectItem is a struct that represents a project item.
type ProjectItem struct {
	URL string
}

// ProjectV2Setter is an interface for setting project items.
type ProjectV2Setter interface {
	Set(categoryID string, statusID string, projectItems []ProjectItem) error
}
