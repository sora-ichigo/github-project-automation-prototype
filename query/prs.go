package query

import "github.com/igsr5/github-project-automation/usecase"

type prFetcherImpl struct{}

// NewPrFetcher is a factory method to create a new instance of PrFetcher.
func NewPrFetcher() usecase.PrFetcher {
	return &prFetcherImpl{}
}

// WorkInProgressPrs returns a list of pull requests that are in progress.
func (f *prFetcherImpl) WorkInProgressPrs() ([]usecase.PullRequest, error) {
	prs := []usecase.PullRequest{
		{
			URL: "",
		},
	}

	return prs, nil
}

// ReviewRequestedPrs returns a list of pull requests that are requested to review.
func (f *prFetcherImpl) ReviewRequestedPrs() ([]usecase.PullRequest, error) {
	prs := []usecase.PullRequest{
		{
			URL: "",
		},
	}

	return prs, nil
}

// ApprovedPrs returns a list of pull requests that are approved.
func (f *prFetcherImpl) ApprovedPrs() ([]usecase.PullRequest, error) {
	prs := []usecase.PullRequest{
		{
			URL: "",
		},
	}

	return prs, nil
}
