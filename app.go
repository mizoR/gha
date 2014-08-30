package gha

import (
	"log"
	"net/http"
	"os"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
)

type App struct {
	client *github.Client
}

func NewApp() *App {
	return &App{}
}

func (app App) Execute() error {
	var command Command
	var command_name string
	var args []string
	var err error

	if len(os.Args) > 1 {
		command_name = os.Args[1]
	}

	if len(os.Args) > 2 {
		args = os.Args[2:]
	} else {
		args = []string{}
	}

	factory := NewCommandFactory()
	command, err = factory.App(&app).Command(command_name).Create(args)
	if err != nil {
		log.Fatal(err)
	}

	err = command.Execute()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (app App) Client() *github.Client {
	var client *http.Client = nil

	if app.client == nil {
		var t *oauth.Transport = nil

		if token := os.Getenv("GITHUB_ACCESS_TOKEN"); token != "" {
			t = &oauth.Transport{
				Token: &oauth.Token{AccessToken: token},
			}
			client = t.Client()
		}

		app.client = github.NewClient(client)
	}

	return app.client
}
