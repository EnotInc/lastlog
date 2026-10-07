package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
)

func main() {

	amount, err := parseAmount()
	if err != nil {
		io.WriteString(os.Stderr, fmt.Sprintf(" [\033[31m %s \033[0m]\n", err.Error()))
	}

	cmd := exec.Command("git", "--no-pager", "log", "--oneline", "--graph", "--all", amount)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		io.WriteString(os.Stderr, err.Error())
		return
	}
}

func parseAmount() (string, error) {
	const defaultAmount = "-25"

	var input string
	switch len(os.Args) {
	case 1:
		return defaultAmount, nil
	case 2:
		input = os.Args[1]
	default:
		return defaultAmount, fmt.Errorf("wrong amount of arguments!")
	}

	n, err := strconv.Atoi(input)
	if err != nil {
		return defaultAmount, fmt.Errorf("unable to parse input! Used default amount of log lines")
	}

	if n <= 0 {
		return defaultAmount, fmt.Errorf("incorrect input. Used default amount instead")
	}

	return fmt.Sprintf("-%d", n), nil
}
