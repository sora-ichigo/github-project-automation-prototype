package query

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/sora-ichigo/github-project-automation/usecase"
)

type reviewPrFetcherImpl struct{}

// NewReviewPrFetcher is a factory method to create a new instance of ReviewPrFetcher.
func NewReviewPrFetcher() usecase.ReviewPrFetcher {
	return &reviewPrFetcherImpl{}
}

// UnReviewedPrs returns a list of pull requests that are not reviewed.
func (f *reviewPrFetcherImpl) UnReviewedPrs() ([]usecase.PullRequest, error) {
	res, err := searchUnReviewedPRCommand()
	if err != nil {
		return nil, fmt.Errorf("failed to search unreviewed prs: %w", err)
	}

	var prs []usecase.PullRequest
	for _, item := range res.Items {
		prs = append(prs, usecase.PullRequest{URL: item.HTMLURL})
	}

	return prs, nil
}

// CommentedPrs returns a list of pull requests that are commented.
func (f *reviewPrFetcherImpl) ReviewedPrs() ([]usecase.PullRequest, error) {
	res, err := searchCommentedReviewPRCommand()
	if err != nil {
		return nil, fmt.Errorf("failed to search commented prs: %w", err)
	}

	var prs []usecase.PullRequest
	for _, item := range res.Items {
		prs = append(prs, usecase.PullRequest{URL: item.HTMLURL})
	}

	return prs, nil
}

// ApprovedPrs returns a list of pull requests that are approved.
func (f *reviewPrFetcherImpl) ApprovedPrs() ([]usecase.PullRequest, error) {
	res, err := searchApprovedReviewPRCommand()
	if err != nil {
		return nil, fmt.Errorf("failed to search approved prs: %w", err)
	}

	var prs []usecase.PullRequest
	for _, item := range res.Items {
		prs = append(prs, usecase.PullRequest{URL: item.HTMLURL})
	}

	return prs, nil
}

// Search query:
// - owner: wantedly
// - state: open
// - review: none
// - review-requested: sora-ichigo
func searchUnReviewedPRCommand() (*SearchQueryResponse, error) {
	cmd := exec.Command("gh", "api", "/search/issues?q=is:open+owner:wantedly+type:pr+review-requested:sora-ichigo")
	output, err := cmd.Output()

	var res SearchQueryResponse
	if err := json.Unmarshal(output, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal unreviewed prs: %w", err)
	}

	return &res, err
}

// Search query:
// - owner: wantedly
// - state: open
// - reviewed-by: sora-ichigo
// - -assignee: sora-ichigo
// - -review: approved
// - -review-requested: sora-ichigo
func searchCommentedReviewPRCommand() (*SearchQueryResponse, error) {
	cmd := exec.Command("gh", "api", "/search/issues?q=is:open+owner:wantedly+type:pr+reviewed-by:sora-ichigo+-assignee:sora-ichigo+-review:approved+-review-requested:sora-ichigo")
	output, err := cmd.Output()

	var res SearchQueryResponse
	if err := json.Unmarshal(output, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal commented prs: %w", err)
	}

	return &res, err
}

// Search query:
// - owner: wantedly
// - state: open
// - review: approved
// - reviewed-by: @me
func searchApprovedReviewPRCommand() (*SearchQueryResponse, error) {
	cmd := exec.Command("gh", "api", "/search/issues?q=is:open+owner:wantedly+type:pr+reviewed-by:sora-ichigo+-assignee:sora-ichigo+review:approved")

	output, err := cmd.Output()

	var res SearchQueryResponse
	if err := json.Unmarshal(output, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal approved prs: %w", err)
	}

	return &res, err
}
