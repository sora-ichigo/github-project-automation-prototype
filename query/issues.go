package query

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/igsr5/github-project-automation/usecase"
)

type issueFetcherImpl struct {
}

// NewIssueFetcher is a factory method to create a new instance of IssueFetcher.
func NewIssueFetcher() usecase.IssueFetcher {
	return &issueFetcherImpl{}
}

// MyIssues returns issues assigned to me.
func (f *issueFetcherImpl) MyIssues() ([]usecase.Issue, error) {
	b, err := searchIssueCommand()
	if err != nil {
		return nil, fmt.Errorf("failed to search issues: %w", err)
	}

	var issues []usecase.Issue
	if err := json.Unmarshal(b, &issues); err != nil {
		return nil, fmt.Errorf("failed to unmarshal issues: %w", err)
	}

	return issues, nil
}

// Search query:
// - owner: wantedly
// - assignee: @me
// - state: open
func searchIssueCommand() ([]byte, error) {
	cmd := exec.Command("gh", "search", "issues", "--owner", "wantedly", "--assignee", "@me", "--state", "open", "--limit", "100", "--json", "url")
	output, err := cmd.Output()
	return output, err
}
