package query

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/sora-ichigo/github-project-automation/usecase"
)

type issueFetcherImpl struct {
}

// NewIssueFetcher is a factory method to create a new instance of IssueFetcher.
func NewIssueFetcher() usecase.IssueFetcher {
	return &issueFetcherImpl{}
}

// MyIssues returns issues assigned to me.
func (f *issueFetcherImpl) MyIssues() ([]usecase.Issue, error) {
	res, err := searchIssueCommand()
	if err != nil {
		return nil, fmt.Errorf("failed to search issues: %w", err)
	}

	var issues []usecase.Issue
	for _, item := range res.Items {
		issues = append(issues, usecase.Issue{URL: item.HTMLURL})
	}

	return issues, nil
}

// Search query:
// - assignee:sora-ichigo
// - is:open
// - owner:wantedly
// - type:issue
func searchIssueCommand() (*SearchQueryResponse, error) {
	cmd := exec.Command("gh", "api", "/search/issues?q=assignee:@me+is:open+owner:wantedly+type:issue")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute gh api command: %w", err)
	}

	var res SearchQueryResponse
	if err := json.Unmarshal(output, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal issues: %w", err)
	}

	return &res, err
}
