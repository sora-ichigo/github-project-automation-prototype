package command

import (
	"testing"

	"github.com/igsr5/github-project-automation/domain"
	mock_domain "github.com/igsr5/github-project-automation/domain/mock"
	gomock "go.uber.org/mock/gomock"
)

func TestIssueAssignedToMeAutomationImpl_MyIssues(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mocks
	mockIssueFetcher := mock_domain.NewMockIssueFetcher(ctrl)
	mockProjectV2Setter := mock_domain.NewMockProjectV2Setter(ctrl)

	// Expected data
	expectedIssues := []domain.Issue{
		{Url: "https://github.com/wantedly/wantedly/issue/1234"},
	}

	// Expected calls
	mockIssueFetcher.EXPECT().MyIssues().Return(expectedIssues, nil).Times(1)
	mockProjectV2Setter.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)

	// Initialize the struct with mocks
	automation := NewIssueAssignedToMeAutomation(mockIssueFetcher, mockProjectV2Setter)

	// Execute the method
	err := automation.SetInProgress()

	// Validate
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
