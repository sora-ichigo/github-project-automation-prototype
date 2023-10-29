package query

import (
	"encoding/json"
	"fmt"
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
		return nil, fmt.Errorf("failed to search unreviewed prs: %w", err)
	}

	var prs []usecase.PullRequest
	if err := json.Unmarshal(b, &prs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal unreviewed prs: %w", err)
	}

	return prs, nil
}

// CommentedPrs returns a list of pull requests that are commented.
func (f *reviewPrFetcherImpl) ReviewedPrs() ([]usecase.PullRequest, error) {
	b, err := searchCommentedReviewPRCommand()
	if err != nil {
		return nil, fmt.Errorf("failed to search commented prs: %w", err)
	}

	var prs []usecase.PullRequest
	if err := json.Unmarshal(b, &prs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal commented prs: %w", err)
	}

	// TODO: review-requested な PR が含まれている場合は除外する

	return prs, nil
}

// ApprovedPrs returns a list of pull requests that are approved.
func (f *reviewPrFetcherImpl) ApprovedPrs() ([]usecase.PullRequest, error) {
	b, err := searchApprovedReviewPRCommand()
	if err != nil {
		return nil, fmt.Errorf("failed to search approved prs: %w", err)
	}

	var prs []usecase.PullRequest
	if err := json.Unmarshal(b, &prs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal approved prs: %w", err)
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
// - review: none
// - reviewed-by: @me
// - no-assignee: @me
func searchCommentedReviewPRCommand() ([]byte, error) {
	cmd := exec.Command("gh", "search", "prs", "--owner", "wantedly", "--state", "open", "--reviewed-by", "@me", "--review", "none", "--no-assignee", "@me", "--limit", "100", "--json", "url")
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
