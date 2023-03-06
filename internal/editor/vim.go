package editor

import (
	"os"
	"os/exec"
)

func OpenVim(args ...string) error {
	c := exec.Command("vim", args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
