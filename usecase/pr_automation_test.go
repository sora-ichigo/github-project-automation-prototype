package usecase_test

import (
	"testing"

	usecase "github.com/igsr5/github-project-automation/usecase"
	mock_usecase "github.com/igsr5/github-project-automation/usecase/mock"
	"go.uber.org/mock/gomock"
)

func TestPrAutomationSetInProgress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mocks
	mockPrFetcher := mock_usecase.NewMockPrFetcher(ctrl)
	mockProjectV2Setter := mock_usecase.NewMockProjectV2Setter(ctrl)

	// Expected data
	expectedIssues := []usecase.PullRequest{
		{URL: "https://github.com/wantedly/wantedly/issue/1234"},
	}

	// Expected calls
	mockPrFetcher.EXPECT().WorkInProgressPrs().Return(expectedIssues, nil).Times(1)
	mockProjectV2Setter.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)

	// Initialize the struct with mocks
	automation := usecase.NewPrAutomation(mockPrFetcher, mockProjectV2Setter)

	// Execute the method
	err := automation.SetInProgress()

	// Validate
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

}

func TestPrAutomationSetInPending(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mocks
	mockPrFetcher := mock_usecase.NewMockPrFetcher(ctrl)
	mockProjectV2Setter := mock_usecase.NewMockProjectV2Setter(ctrl)

	// Expected data
	expectedIssues := []usecase.PullRequest{
		{URL: "https://github.com/wantedly/wantedly/issue/1234"},
	}

	// Expected calls
	mockPrFetcher.EXPECT().ReviewRequestedPrs().Return(expectedIssues, nil).Times(1)
	mockProjectV2Setter.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)

	// Initialize the struct with mocks
	automation := usecase.NewPrAutomation(mockPrFetcher, mockProjectV2Setter)

	// Execute the method
	err := automation.SetInPending()

	// Validate
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

}

func TestPrAutomationSetComplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mocks
	mockPrFetcher := mock_usecase.NewMockPrFetcher(ctrl)
	mockProjectV2Setter := mock_usecase.NewMockProjectV2Setter(ctrl)

	// Expected data
	expectedIssues := []usecase.PullRequest{
		{URL: "https://github.com/wantedly/wantedly/issue/1234"},
	}

	// Expected calls
	mockPrFetcher.EXPECT().ApprovedPrs().Return(expectedIssues, nil).Times(1)
	mockProjectV2Setter.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)

	// Initialize the struct with mocks
	automation := usecase.NewPrAutomation(mockPrFetcher, mockProjectV2Setter)

	// Execute the method
	err := automation.SetComplete()

	// Validate
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

}
