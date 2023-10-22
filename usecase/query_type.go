//go:generate mockgen -source=query_type.go -destination=mock/query_type_mock.go -package=mock

package usecase

// Issue is a struct for an issue.
type Issue struct {
	URL string
}

// PullRequest is a struct for a pull request.
type PullRequest struct {
	URL string
}

// IssueFetcher is an interface for fetching issues.
type IssueFetcher interface {
	MyIssues() ([]Issue, error)
}

// PrFetcher is an interface for fetching pull requests.
type PrFetcher interface {
	WorkInProgressPrs() ([]PullRequest, error)
	ReviewRequestedPrs() ([]PullRequest, error)
	ApprovedPrs() ([]PullRequest, error)
}
