package query

import (
	"encoding/json"
	"errors"
	"os/exec"

	"github.com/igsr5/github-project-automation/usecase"
)

type reviewPrFetcherImpl struct{}

// NewReviewPrFetcher is a factory method to create a new instance of ReviewPrFetcher.
func NewReviewPrFetcher() usecase.ReviewPrFetcher {
	return &reviewPrFetcherImpl{}
}

// UnReviewedPrs returns a list of pull requests that are not reviewed.
func (f *reviewPrFetcherImpl) UnReviewedPrs() ([]usecase.PullRequest, error) {
	b, err := searchUnReviewedPRCommand()
	if err != nil {
		return nil, errors.Join(err)
	}

	var prs []usecase.PullRequest
	if err := json.Unmarshal(b, &prs); err != nil {
		return nil, errors.Join(err)
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
	b, err := searchApprovedReviewPRCommand()
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
// - state: open
// - review-requested: @me
func searchUnReviewedPRCommand() ([]byte, error) {
	cmd := exec.Command("gh", "search", "prs", "--owner", "wantedly", "--state", "open", "--review-requested", "@me", "--limit", "100", "--json", "url")
	output, err := cmd.Output()
	return output, err
}

// Search query:
// - owner: wantedly
// - state: open
// - review: approved
// - reviewed-by: @me
func searchApprovedReviewPRCommand() ([]byte, error) {
	cmd := exec.Command("gh", "search", "prs", "--owner", "wantedly", "--state", "open", "--reviewed-by", "@me", "--review", "approved", "--limit", "100", "--json", "url")
	output, err := cmd.Output()
	return output, err
}
