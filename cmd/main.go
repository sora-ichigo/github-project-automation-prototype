package main

import (
	"log"
	"os"

	"github.com/igsr5/github-project-automation/command"
	"github.com/igsr5/github-project-automation/domain"
	"github.com/igsr5/github-project-automation/query"
	"github.com/igsr5/github-project-automation/usecase"
)

func main() {
	projectV2Setter := command.NewProjectV2Setter()
	issueFetcher := query.NewIssueFetcher()
	prFetcher := query.NewPrFetcher()
	reviewPrFetcher := query.NewReviewPrFetcher()
	automations := []domain.Automation{
		usecase.NewIssueAutomation(issueFetcher, projectV2Setter),
		usecase.NewPrAutomation(prFetcher, projectV2Setter),
		usecase.NewReviewPrAutomation(reviewPrFetcher, projectV2Setter),
	}

	for _, a := range automations {
		if err := a.SetInPending(); err != nil {
			log.Printf("error occurred: %v", err)
			os.Exit(1)
		}

		if err := a.SetInPending(); err != nil {
			log.Printf("error occurred: %v", err)
			os.Exit(1)
		}

		if err := a.SetComplete(); err != nil {
			log.Printf("error occurred: %v", err)
			os.Exit(1)
		}
	}
}
