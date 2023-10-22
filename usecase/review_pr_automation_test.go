package usecase_test

import (
	"testing"

	variables "github.com/igsr5/github-project-automation"
	usecase "github.com/igsr5/github-project-automation/usecase"
	mock_usecase "github.com/igsr5/github-project-automation/usecase/mock"
	"go.uber.org/mock/gomock"
)

func TestReviewPrAutomationSetInProgress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mocks
	mockPrFetcher := mock_usecase.NewMockReviewPrFetcher(ctrl)
	mockProjectV2Setter := mock_usecase.NewMockProjectV2Setter(ctrl)

	// Expected data
	expectedPRs := []usecase.PullRequest{
		{URL: "https://github.com/wantedly/wantedly/issue/1234"},
	}

	// Expected calls
	mockPrFetcher.EXPECT().UnReviewedPrs().Return(expectedPRs, nil).Times(1)
	mockProjectV2Setter.EXPECT().Set(variables.REVIEW_PR_CATEGORY_ID, variables.IN_PROGRESS_STATUS_ID, gomock.Any()).Return(nil).Times(1)

	// Initialize the struct with mocks
	automation := usecase.NewReviewPrAutomation(mockPrFetcher, mockProjectV2Setter)

	// Execute the method
	err := automation.SetInProgress()

	// Validate
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

}

func TestReviewPrAutomationSetInPending(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mocks
	mockPrFetcher := mock_usecase.NewMockReviewPrFetcher(ctrl)
	mockProjectV2Setter := mock_usecase.NewMockProjectV2Setter(ctrl)

	// Expected data
	expectedPRs := []usecase.PullRequest{
		{URL: "https://github.com/wantedly/wantedly/issue/1234"},
	}

	// Expected calls
	mockPrFetcher.EXPECT().CommentedPrs().Return(expectedPRs, nil).Times(1)
	mockPrFetcher.EXPECT().ChangesRequestedPrs().Return(expectedPRs, nil).Times(1)
	mockProjectV2Setter.EXPECT().Set(variables.REVIEW_PR_CATEGORY_ID, variables.IN_PENDING_STATUS_ID, gomock.Any()).Return(nil).Times(1)

	// Initialize the struct with mocks
	automation := usecase.NewReviewPrAutomation(mockPrFetcher, mockProjectV2Setter)

	// Execute the method
	err := automation.SetInPending()

	// Validate
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

}

func TestReviewPrAutomationSetComplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mocks
	mockReviewPrFetcher := mock_usecase.NewMockReviewPrFetcher(ctrl)
	mockProjectV2Setter := mock_usecase.NewMockProjectV2Setter(ctrl)

	// Expected data
	expectedPRs := []usecase.PullRequest{
		{URL: "https://github.com/wantedly/wantedly/issue/1234"},
	}

	// Expected calls
	mockReviewPrFetcher.EXPECT().ApprovedPrs().Return(expectedPRs, nil).Times(1)
	mockProjectV2Setter.EXPECT().Set(variables.REVIEW_PR_CATEGORY_ID, variables.COMPLETE_STATUS_ID, gomock.Any()).Return(nil).Times(1)

	// Initialize the struct with mocks
	automation := usecase.NewReviewPrAutomation(mockReviewPrFetcher, mockProjectV2Setter)

	// Execute the method
	err := automation.SetComplete()

	// Validate
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

}
