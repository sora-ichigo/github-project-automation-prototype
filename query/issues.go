package query

import (
	"github.com/igsr5/github-project-automation/usecase"
)

type issueFetcherImpl struct {
}

// NewIssueFetcher is a factory method to create a new instance of IssueFetcher.
func NewIssueFetcher() usecase.IssueFetcher {
	return &issueFetcherImpl{}
}

// MyIssues returns issues assigned to me.
func (f *issueFetcherImpl) MyIssues() ([]usecase.Issue, error) {
	issues := []usecase.Issue{
		{
			Url: "https://github.com/wantedly/wantedly/issue/1234",
		},
	}
	return issues, nil
}
