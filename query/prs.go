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

// UnReviewedPrs returns a list of pull requests that are not reviewed.
func (f *prFetcherImpl) UnReviewedPrs() ([]usecase.PullRequest, error) {
	b, err := searchDraftPRsCommand()
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
	b, err := searchChangeRequestedPRsCommand()
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
func (f *prFetcherImpl) ReviewRequestedPrs() ([]usecase.PullRequest, error) {
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
// - draft: 1
func searchDraftPRsCommand() ([]byte, error) {
	cmd := exec.Command("gh", "search", "prs", "--owner", "wantedly", "--assignee", "@me", "--state", "open", "--draft", "1", "--limit", "100", "--json", "url")
	output, err := cmd.Output()
	return output, err
}

// Search query:
// - owner: wantedly
// - assignee: @me
// - state: open
// - review: changes_requested
func searchChangeRequestedPRsCommand() ([]byte, error) {
	cmd := exec.Command("gh", "search", "prs", "--owner", "wantedly", "--assignee", "@me", "--state", "open", "--review", "changes_requested", "--limit", "100", "--json", "url")
	output, err := cmd.Output()
	return output, err
}

// Search query:
// - owner: wantedly
// - assignee: @me
// - state: open
// - review: none
// - draft: 0
func searchReviewRequestedPRsCommand() ([]byte, error) {
	cmd := exec.Command("gh", "search", "prs", "--owner", "wantedly", "--assignee", "@me", "--state", "open", "--draft", "0", "--review", "none", "--limit", "100", "--json", "url")
	output, err := cmd.Output()
	return output, err
}

// Search query: Approved のレビューがあるPR
// - owner: wantedly
// - assignee: @me
// - state: open
// - review: approved
func searchApprovedPRsCommand() ([]byte, error) {
	cmd := exec.Command("gh", "search", "prs", "--owner", "wantedly", "--assignee", "@me", "--state", "open", "--review", "approved", "--limit", "100", "--json", "url")
	output, err := cmd.Output()
	return output, err
}
