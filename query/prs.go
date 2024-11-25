package query

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/sora-ichigo/github-project-automation/usecase"
)

type prFetcherImpl struct{}

// NewPrFetcher is a factory method to create a new instance of PrFetcher.
func NewPrFetcher() usecase.PrFetcher {
	return &prFetcherImpl{}
}

// UnReviewedPrs returns a list of pull requests that are not reviewed.
func (f *prFetcherImpl) UnReviewedPrs() ([]usecase.PullRequest, error) {
	res, err := searchUnReviewedPRsCommand()
	if err != nil {
		return nil, fmt.Errorf("failed to search draft prs: %w", err)
	}

	var prs []usecase.PullRequest
	for _, item := range res.Items {
		prs = append(prs, usecase.PullRequest{URL: item.HTMLURL})
	}

	return prs, nil
}

// ChangesRequestedPrs returns a list of pull requests that are requested to change.
func (f *prFetcherImpl) ReviewedPrs() ([]usecase.PullRequest, error) {
	res, err := searchReviewedPRsCommand()
	if err != nil {
		return nil, fmt.Errorf("failed to search change requested prs: %w", err)
	}

	var prs []usecase.PullRequest
	for _, item := range res.Items {
		prs = append(prs, usecase.PullRequest{URL: item.HTMLURL})
	}

	return prs, nil
}

// ApprovedPrs returns a list of pull requests that are approved.
func (f *prFetcherImpl) ApprovedPrs() ([]usecase.PullRequest, error) {
	res, err := searchApprovedPRsCommand()
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
// - assignee:sora-ichigo
// - is:open
// - owner:wantedly
// - type:pr
// - review:none
// - draft:true
func searchUnReviewedPRsCommand() (*SearchQueryResponse, error) {
	// gh api "/search/issues?q=assignee:sora-ichigo+is:open+owner:wantedly+type:pr+review:none"
	cmd := exec.Command("gh", "api", "/search/issues?q=assignee:sora-ichigo+is:open+owner:wantedly+type:pr+review:none+draft:true")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute gh api command: %w", err)
	}

	var res SearchQueryResponse
	if err := json.Unmarshal(output, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal unreviewed prs: %w", err)
	}
	return &res, nil
}

// Search query:
// - assignee:sora-ichigo
// - is:open
// - owner:wantedly
// - type:pr
// - -review:approved
// - draft:false
func searchReviewedPRsCommand() (*SearchQueryResponse, error) {
	cmd := exec.Command("gh", "api", "/search/issues?q=assignee:sora-ichigo+is:open+owner:wantedly+type:pr+draft:false+-review:approved")
	output, err := cmd.Output()

	var res SearchQueryResponse
	if err := json.Unmarshal(output, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal reviewed prs: %w", err)
	}

	return &res, err
}

// Search query: Approved のレビューがあるPR
// - assignee:sora-ichigo
// - is:open
// - owner:wantedly
// - type:pr
// - review:approved
func searchApprovedPRsCommand() (*SearchQueryResponse, error) {
	// gh api "/search/issues?q=assignee:sora-ichigo+is:open+owner:wantedly+type:pr+review:approved"
	cmd := exec.Command("gh", "api", "/search/issues?q=assignee:sora-ichigo+is:open+owner:wantedly+type:pr+review:approved")
	output, err := cmd.Output()

	var res SearchQueryResponse
	if err := json.Unmarshal(output, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal approved prs: %w", err)
	}

	return &res, err
}
