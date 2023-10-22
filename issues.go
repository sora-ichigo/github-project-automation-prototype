package main

// Issue is a struct for an issue.
type Issue struct {
	Url string
}

// IssueFetcher is an interface for fetching issues.
type IssueFetcher interface {
	MyIssues() ([]Issue, error)
}

type issueFetcherImpl struct {
}

// NewIssueFetcher is a factory method to create a new instance of IssueFetcher.
func NewIssueFetcher() IssueFetcher {
	return &issueFetcherImpl{}
}

// MyIssues returns issues assigned to me.
func (f *issueFetcherImpl) MyIssues() ([]Issue, error) {
	issues := []Issue{
		{
			Url: "https://github.com/wantedly/wantedly/issue/1234",
		},
	}
	return issues, nil
}
