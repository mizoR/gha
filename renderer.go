package gha

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/google/go-github/github"
)

type Renderer struct {
}

func (r Renderer) zeropadding(src string, length int) string {
	var i int = 0
	var w int = 0
	var v rune
	dst := []rune{}

	for _, v = range src {
		if w < length {
			dst = append(dst, v)
			if utf8.RuneLen(v) > 1 {
				w += 2
			} else {
				w++
			}
		} else {
			break
		}
	}

	for i = w; i < length; i++ {
		dst = append(dst, ' ')
	}

	return string(dst)
}

type IssueListRenderer struct {
	issues []github.Issue

	Renderer
}

func (r IssueListRenderer) render() {
	for i := 0; i < len(r.issues); i++ {
		r.renderIssue(r.issues[i])
	}
}

func (r IssueListRenderer) renderIssue(issue github.Issue) {
	padding := GetWinsize().Col - 100
	fmt.Printf(r.format(), *issue.Number, *issue.State, r.zeropadding(*issue.Title, padding), *issue.User.Login, (*issue.UpdatedAt).String())
}

func (r IssueListRenderer) format() string {
	values := []string{"#%-5d", "\x1b[36m%6s\x1b[0m", "%s", "\x1b[35m%15s\x1b[0m", "%s"}
	return strings.Join(values, "\t") + "\n"
}

type IssueRenderer struct {
	issue    *github.Issue
	comments []github.IssueComment

	Renderer
}

func (r IssueRenderer) render() {
	padding := GetWinsize().Col - 80
	issue := r.issue
	comments := r.comments

	fmt.Printf(r.format(), *issue.State, *issue.Number, r.zeropadding(*issue.Title, padding), *issue.User.Login, (*issue.CreatedAt).String(), *issue.Body)
	fmt.Printf("\n\n")
	for i := 0; i < len(comments); i++ {
		comment := comments[i]
		fmt.Printf("#%d commented by @%s at %s \n%s \n\n", *comment.ID, *comment.User.Login, *comment.CreatedAt, *comment.Body)
	}
}

func (r IssueRenderer) format() string {
	values := []string{"[%s] \x1b[36m#%d\x1b[0m %s (created this issue by \x1b[35m@%s\x1b[0m at %s)\n\n%s"}
	return strings.Join(values, "\t") + "\n"
}
