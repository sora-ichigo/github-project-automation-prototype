package query

type SearchQueryResponse struct {
	Items []Item `json:"items"`
}

type Item struct {
	HTMLURL string `json:"html_url"`
}
