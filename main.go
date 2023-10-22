package main

import (
	"log"
	"os"
)

func main() {
	projectV2Setter := NewProjectV2Setter()
	issueFetcher := NewIssueFetcher()
	automations := []Automation{
		NewIssueAssignedToMeAutomation(issueFetcher, projectV2Setter),
		NewPrAssignedToMeAutomation(),
		NewPrAssignedToReviewerAutomation(),
	}

	for _, a := range automations {
		if err := a.setInPending(); err != nil {
			log.Printf("error occurred: %v", err)
			os.Exit(1)
		}

		if err := a.setInPending(); err != nil {
			log.Printf("error occurred: %v", err)
			os.Exit(1)
		}

		if err := a.setComplete(); err != nil {
			log.Printf("error occurred: %v", err)
			os.Exit(1)
		}
	}
}
