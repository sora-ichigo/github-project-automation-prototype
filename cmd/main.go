package main

import (
	"fmt"
	"log"
	"os"

	"github.com/igsr5/github-project-automation/command"
	"github.com/igsr5/github-project-automation/domain"
	"github.com/igsr5/github-project-automation/query"
	"github.com/igsr5/github-project-automation/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/run", func(c echo.Context) error {
		err := run()
		if err != nil {
			log.Printf("error: %v", err)
			return c.String(500, "error")
		}

		return c.String(200, "ok")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "ok")
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
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
	return nil
}
