package query

import "github.com/igsr5/github-project-automation/usecase"

type reviewPrFetcherImpl struct{}

// NewReviewPrFetcher is a factory method to create a new instance of ReviewPrFetcher.
func NewReviewPrFetcher() usecase.ReviewPrFetcher {
	return &reviewPrFetcherImpl{}
}

// UnReviewedPrs returns a list of pull requests that are not reviewed.
func (f *reviewPrFetcherImpl) UnReviewedPrs() ([]usecase.PullRequest, error) {
	// TODO: Implement
	prs := []usecase.PullRequest{
		{
			URL: "",
		},
	}
	return prs, nil
}

// CommentedPrs returns a list of pull requests that are commented.
func (f *reviewPrFetcherImpl) CommentedPrs() ([]usecase.PullRequest, error) {
	// TODO: Implement
	prs := []usecase.PullRequest{
		{
			URL: "",
		},
	}
	return prs, nil
}

// ChangesRequestedPrs returns a list of pull requests that are requested to change.
func (f *reviewPrFetcherImpl) ChangesRequestedPrs() ([]usecase.PullRequest, error) {
	// TODO: Implement
	prs := []usecase.PullRequest{
		{
			URL: "",
		},
	}
	return prs, nil
}

// ApprovedPrs returns a list of pull requests that are approved.
func (f *reviewPrFetcherImpl) ApprovedPrs() ([]usecase.PullRequest, error) {
	// TODO: Implement
	prs := []usecase.PullRequest{
		{
			URL: "",
		},
	}
	return prs, nil
}
