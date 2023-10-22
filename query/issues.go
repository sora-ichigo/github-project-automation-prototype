package query

import "github.com/igsr5/github-project-automation/domain"

type issueFetcherImpl struct {
}

// NewIssueFetcher is a factory method to create a new instance of IssueFetcher.
func NewIssueFetcher() domain.IssueFetcher {
	return &issueFetcherImpl{}
}

// MyIssues returns issues assigned to me.
func (f *issueFetcherImpl) MyIssues() ([]domain.Issue, error) {
	issues := []domain.Issue{
		{
			Url: "https://github.com/wantedly/wantedly/issue/1234",
		},
	}
	return issues, nil
}
