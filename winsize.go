package gha

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

type Winsize struct {
	Row int
	Col int
}

var winsize *Winsize = nil

func GetWinsize() *Winsize {
	if winsize == nil {
		var row int
		var col int
		cmd := exec.Command("stty", "size")
		cmd.Stdin = os.Stdin
		out, err := cmd.Output()

		re := regexp.MustCompile("([0-9]+) ([0-9]+)\n")
		matches := re.FindStringSubmatch(string(out))

		if len(matches) > 2 {
			row, _ = strconv.Atoi(matches[1])
			col, _ = strconv.Atoi(matches[2])
		} else {
			fmt.Println("Could not get winsize...")
		}

		if err != nil {
			log.Fatal(err)
		}

		return &Winsize{
			Row: row,
			Col: col,
		}
	}

	return winsize
}
