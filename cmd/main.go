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

	e.POST("/issues", func(c echo.Context) error {
		err := run("issues")
		if err != nil {
			log.Printf("error: %v", err)
			return c.String(500, "error")
		}

		return c.String(200, "ok")

	})

	e.POST("/pull_requests", func(c echo.Context) error {
		err := run("pull_requests")
		if err != nil {
			log.Printf("error: %v", err)
			return c.String(500, "error")
		}

		return c.String(200, "ok")
	})

	e.POST("/review_pull_requests", func(c echo.Context) error {
		err := run("review_pull_requests")
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

func run(mode string) error {
	projectV2Setter := command.NewProjectV2Setter()
	issueFetcher := query.NewIssueFetcher()
	prFetcher := query.NewPrFetcher()
	reviewPrFetcher := query.NewReviewPrFetcher()

	var a domain.Automation

	if mode == "issues" {
		a = usecase.NewIssueAutomation(issueFetcher, projectV2Setter)
	} else if mode == "pull_requests" {
		a = usecase.NewPrAutomation(prFetcher, projectV2Setter)
	} else if mode == "review_pull_requests" {
		a = usecase.NewReviewPrAutomation(reviewPrFetcher, projectV2Setter)
	} else {
		return fmt.Errorf("invalid mode: %s", mode)
	}

	if err := a.SetInProgress(); err != nil {
		return fmt.Errorf("automations is failed: %v", err)
	}

	if err := a.SetInPending(); err != nil {
		return fmt.Errorf("automations is failed: %v", err)
	}

	if err := a.SetComplete(); err != nil {
		return fmt.Errorf("automations is failed: %v", err)
	}

	return nil
}
