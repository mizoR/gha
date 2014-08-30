package gha

import (
	"github.com/google/go-github/github"
	flags "github.com/jessevdk/go-flags"
)

type IssueListCommand struct {
	app  *App
	opts *IssueListCommandOptions
}

type IssueListCommandOptions struct {
	All bool   `long:"all"`
	Org string `long:"org"`

	Filter    string `long:"filter"`
	State     string `long:"state"`
	Sort      string `long:"sort"`
	Direction string `long:"direction"`

	ListOptions
}

func (opts IssueListCommandOptions) Parse() *github.IssueListOptions {
	o := &github.IssueListOptions{
		Filter:    opts.Filter,
		State:     opts.State,
		Sort:      opts.Sort,
		Direction: opts.Direction,
	}

	o.Page = opts.Page
	o.PerPage = opts.PerPage

	return o
}

func NewIssueListCommand(app *App, args []string) (*IssueListCommand, error) {
	var opts IssueListCommandOptions

	args, err := flags.ParseArgs(&opts, args)
	if err != nil {
		return nil, err
	}

	return &IssueListCommand{app: app, opts: &opts}, nil
}

func (list IssueListCommand) Execute() error {
	var issues []github.Issue
	var err error
	opts := list.opts.Parse()

	if list.opts.Org != "" {
		issues, _, err = list.app.Client().Issues.ListByOrg(list.opts.Org, opts)
	} else {
		issues, _, err = list.app.Client().Issues.List(list.opts.All, opts)
	}
	if err != nil {
		return err
	}

	renderer := IssueListRenderer{issues: issues}
	renderer.render()

	return nil
}
