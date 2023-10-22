package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_main "github.com/igsr5/github-project-automation/mock"
)

func TestIssueAssignedToMeAutomationImpl_MyIssues(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mocks
	mockIssueFetcher := mock_main.MockIssueFetcher(ctrl)
	mockProjectV2Setter := mock_main.MockProjectV2Setter(ctrl)

	// Expected data
	expectedIssues := []Issue{
		{Url: "https://github.com/wantedly/wantedly/issue/1234"},
	}

	// Expected calls
	mockIssueFetcher.EXPECT().MyIssues().Return(expectedIssues, nil).Times(1)
	mockProjectV2Setter.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)

	// Initialize the struct with mocks
	automation := NewIssueAssignedToMeAutomation(mockIssueFetcher, mockProjectV2Setter)

	// Execute the method
	err := automation.setInProgress()

	// Validate
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
