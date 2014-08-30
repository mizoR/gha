package gha

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
)

type GitConfig struct {
	Repo string
}

var gitconfig *GitConfig

func GetGitConfig() *GitConfig {
	if gitconfig == nil {
		var stdout bytes.Buffer
		var repo string

		// Challenge to get params by `git config`
		cmd := exec.Command("git", "config", "remote.origin.url")
		cmd.Stdout = &stdout

		if err := cmd.Run(); err != nil {
			fmt.Println(err)
		}

		re := regexp.MustCompile("github.com:(.+).git")
		matches := re.FindStringSubmatch(stdout.String())

		if len(matches) > 1 {
			repo = matches[1]
		} else {
			fmt.Println("Repository is not found...")
		}

		gitconfig = &GitConfig{
			Repo: repo,
		}
	}

	return gitconfig
}
