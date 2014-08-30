package gha

import "fmt"

type HelpCommand struct {
}

func NewHelpCommand(_ *App, args []string) (*HelpCommand, error) {
	return &HelpCommand{}, nil
}

func (command HelpCommand) Execute() error {
	// FIXME
	fmt.Println("Github Issue HelpCommand")
	return nil
}
