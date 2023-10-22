//go:generate mockgen -source=issues.go -destination=mock/issues_mock.go -package=mock

package usecase

// Issue is a struct for an issue.
type Issue struct {
	URL string
}

// IssueFetcher is an interface for fetching issues.
type IssueFetcher interface {
	MyIssues() ([]Issue, error)
}
