package variables

var (
	// IN_PROGRESS_STATUS_ID is a status id for "In Progress" column.
	IN_PROGRESS_STATUS_ID = "47fc9ee4"
	// IN_PENDING_STATUS_ID is a status id for "In Pending" column.
	IN_PENDING_STATUS_ID = "e0c049ea"
	// COMPLETE_STATUS_ID is a status id for "Complete" column.
	COMPLETE_STATUS_ID = "f2f3bb33"
)

var (
	// PROJECT_ID is a project id.
	PROJECT_ID = "227"
)

var (
	// ISSUE_CATEGORY_ID is a category id for issues.
	ISSUE_CATEGORY_ID = "b41c9a4b"
	// PR_CATEGORY_ID is a category id for review.
	PR_CATEGORY_ID = "086cc5ba"
	// REVIEW_PR_CATEGORY_ID a status id for "In Progress" column.
	REVIEW_PR_CATEGORY_ID = "5f8f94e2"
)

// NOTE: this command can investigate the project data.
// ```
// $ gh project view 227 --owner wantedly --format json | jq
// $ gh project field-list 227 --owner wantedly --format json | jq
// ```
