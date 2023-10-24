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
	err := run()
	if err != nil {
		log.Printf("error: %v", err)
	}
}

func run() error {
	projectV2Setter := command.NewProjectV2Setter()
	issueFetcher := query.NewIssueFetcher()
	prFetcher := query.NewPrFetcher()
	reviewPrFetcher := query.NewReviewPrFetcher()
	automations := []domain.Automation{
		usecase.NewReviewPrAutomation(reviewPrFetcher, projectV2Setter),
		usecase.NewPrAutomation(prFetcher, projectV2Setter),
		usecase.NewIssueAutomation(issueFetcher, projectV2Setter),
	}

	for _, a := range automations {
		if err := a.SetInProgress(); err != nil {
			return fmt.Errorf("automations is failed: %v", err)
		}

		if err := a.SetInPending(); err != nil {
			return fmt.Errorf("automations is failed: %v", err)
		}

		if err := a.SetComplete(); err != nil {
			return fmt.Errorf("automations is failed: %v", err)
		}
	}
}
