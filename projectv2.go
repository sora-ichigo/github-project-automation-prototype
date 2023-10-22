package main

type ProjectV2Setter interface {
	Set(categoryId string, statusId string) error
}

type ProjectV2SetterImpl struct {
}

func NewProjectV2Setter() ProjectV2Setter {
	return &ProjectV2SetterImpl{}
}

func (s *ProjectV2SetterImpl) Set(categoryId string, statusId string) error {
	return nil
}
