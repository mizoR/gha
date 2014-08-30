package gha

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/google/go-github/github"
)

type IssueCommand struct {
	app    *App
	owner  string
	repo   string
	number int
}

func NewIssueCommand(app *App, args []string) (*IssueCommand, error) {
	if len(args) != 3 {
		msg := fmt.Sprintf("ArgumentError: wrong number of argument(%d for 3)", len(args))
		return nil, errors.New(msg)
	}

	owner := args[0]
	repo := args[1]
	number, err := strconv.Atoi(args[2])
	if err != nil {
		return nil, err
	}

	return &IssueCommand{app: app, owner: owner, repo: repo, number: number}, nil
}

func (c IssueCommand) Execute() error {
	var issue *github.Issue
	var err error

	issue, _, err = c.app.Client().Issues.Get(c.owner, c.repo, c.number)
	if err != nil {
		return err
	}

	renderer := IssueRenderer{issue: issue}
	renderer.render()

	return nil
}
