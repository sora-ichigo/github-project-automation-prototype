package main

import (
	"log"
	"os"

	"github.com/igsr5/github-project-automation/command"
	"github.com/igsr5/github-project-automation/domain"
)

func main() {
	projectV2Setter := command.NewProjectV2Setter()
	issueFetcher := command.NewIssueFetcher()
	automations := []domain.Automation{
		command.NewIssueAssignedToMeAutomation(issueFetcher, projectV2Setter),
		command.NewPrAssignedToMeAutomation(),
		command.NewPrAssignedToReviewerAutomation(),
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
