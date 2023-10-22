package query

import (
	"encoding/json"
	"errors"
	"os/exec"

	"github.com/igsr5/github-project-automation/usecase"
)

type prFetcherImpl struct{}

// NewPrFetcher is a factory method to create a new instance of PrFetcher.
func NewPrFetcher() usecase.PrFetcher {
	return &prFetcherImpl{}
}

// WorkInProgressPrs returns a list of pull requests that are in progress.
func (f *prFetcherImpl) UnReviewedPrs() ([]usecase.PullRequest, error) {
	b, err := searchUnReviewRequestedPRsCommand()
	if err != nil {
		return nil, errors.Join(err)
	}

	var prs []usecase.PullRequest
	if err := json.Unmarshal(b, &prs); err != nil {
		return nil, errors.Join(err)
	}

	return prs, nil
}

// ReviewRequestedPrs returns a list of pull requests that are requested to review.
func (f *prFetcherImpl) CommentedPrs() ([]usecase.PullRequest, error) {
	b, err := searchReviewRequestedPRsCommand()
	if err != nil {
		return nil, errors.Join(err)
	}

	var prs []usecase.PullRequest
	if err := json.Unmarshal(b, &prs); err != nil {
		return nil, errors.Join(err)
	}

	return prs, nil
}

// ChangesRequestedPrs returns a list of pull requests that are requested to change.
func (f *prFetcherImpl) ChangesRequestedPrs() ([]usecase.PullRequest, error) {
	b, err := searchReviewRequestedPRsCommand()
	if err != nil {
		return nil, errors.Join(err)
	}

	var prs []usecase.PullRequest
	if err := json.Unmarshal(b, &prs); err != nil {
		return nil, errors.Join(err)
	}

	return prs, nil
}

// ApprovedPrs returns a list of pull requests that are approved.
func (f *prFetcherImpl) ApprovedPrs() ([]usecase.PullRequest, error) {
	b, err := searchApprovedPRsCommand()
	if err != nil {
		return nil, errors.Join(err)
	}

	var prs []usecase.PullRequest
	if err := json.Unmarshal(b, &prs); err != nil {
		return nil, errors.Join(err)
	}

	return prs, nil
}

// Search query:
// - owner: wantedly
// - assignee: @me
// - state: open
func searchUnReviewRequestedPRsCommand() ([]byte, error) {
	// TODO: Implement
	cmd := exec.Command("gh", "search", "issues", "--owner", "wantedly", "--assignee", "@me", "--state", "open", "--limit", "100", "--json", "url")
	output, err := cmd.Output()
	return output, err
}

// Search query:
// - owner: wantedly
// - assignee: @me
// - state: open
func searchReviewRequestedPRsCommand() ([]byte, error) {
	// TODO: Implement
	cmd := exec.Command("gh", "search", "issues", "--owner", "wantedly", "--assignee", "@me", "--state", "open", "--limit", "100", "--json", "url")
	output, err := cmd.Output()
	return output, err
}

// Search query:
// - owner: wantedly
// - assignee: @me
// - state: open
// - review: approved
func searchApprovedPRsCommand() ([]byte, error) {
	cmd := exec.Command("gh", "search", "prs", "--owner", "wantedly", "--assignee", "@me", "--state", "open", "--review", "approved", "--limit", "100", "--json", "url")
	output, err := cmd.Output()
	return output, err
}
