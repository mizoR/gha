package gha

type CommandFactory struct {
	app     *App
	command string
}

func NewCommandFactory() *CommandFactory {
	return &CommandFactory{command: "help"}
}

func (factory *CommandFactory) App(app *App) *CommandFactory {
	factory.app = app

	return factory
}

func (factory *CommandFactory) Command(command_name string) *CommandFactory {
	factory.command = command_name

	return factory
}

func (factory CommandFactory) Create(args []string) (Command, error) {
	if factory.command == "list" {
		return NewIssueListCommand(factory.app, args)
	}

	if factory.command == "show" {
		return NewIssueCommand(factory.app, args)
	}

	return NewHelpCommand(factory.app, args)
}
