//go:generate mockgen -source=query_type.go -destination=mock/query_type_mock.go -package=mock

package usecase

// Issue is a struct for an issue.
type Issue struct {
	URL string `json:"url"`
}

// PullRequest is a struct for a pull request.
type PullRequest struct {
	URL string `json:"url"`
}

// IssueFetcher is an interface for fetching issues.
type IssueFetcher interface {
	MyIssues() ([]Issue, error)
}

// PrFetcher is an interface for fetching pull requests.
type PrFetcher interface {
	UnReviewedPrs() ([]PullRequest, error)
	CommentedPrs() ([]PullRequest, error)
	ChangesRequestedPrs() ([]PullRequest, error)
	ApprovedPrs() ([]PullRequest, error)
}

// ReviewPrFetcher is an interface for fetching pull requests that are requested to review.
type ReviewPrFetcher interface {
	UnReviewedPrs() ([]PullRequest, error)
	CommentedPrs() ([]PullRequest, error)
	ChangesRequestedPrs() ([]PullRequest, error)
	ApprovedPrs() ([]PullRequest, error)
}
