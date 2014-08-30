package gha

type Command interface {
	Execute() error
}
