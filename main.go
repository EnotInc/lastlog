package main

import (
	"io"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "--no-pager", "log", "--oneline", "--graph", "--all", "-25")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		io.WriteString(os.Stderr, err.Error())
		return
	}
}
