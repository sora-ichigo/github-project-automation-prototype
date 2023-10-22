package main

type Issue struct {
	Url string
}
type IssueFetcher interface {
	MyIssues() ([]Issue, error)
}

type issueFetcherImpl struct {
}

func NewIssueFetcher() IssueFetcher {
	return &issueFetcherImpl{}
}

func (f *issueFetcherImpl) MyIssues() ([]Issue, error) {
	issues := []Issue{
		{
			Url: "https://github.com/wantedly/wantedly/issue/1234",
		},
	}
	return issues, nil
}
