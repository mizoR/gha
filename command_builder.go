package gha

type CommandBuilder struct {
	app     *App
	command string
}

func NewCommandBuilder() *CommandBuilder {
	return &CommandBuilder{command: "help"}
}

func (builder *CommandBuilder) App(app *App) *CommandBuilder {
	builder.app = app

	return builder
}

func (builder *CommandBuilder) Command(command_name string) *CommandBuilder {
	builder.command = command_name

	return builder
}

func (builder CommandBuilder) Build(args []string) (Command, error) {
	if builder.command == "list" {
		return NewIssueListCommand(builder.app, args)
	}

	if builder.command == "show" {
		return NewIssueCommand(builder.app, args)
	}

	return NewHelpCommand(builder.app, args)
}
