package peco

import (
	"fmt"
	"os/exec"
)

func OpenPeco(txbufDir string) (string, error) {
	cmd := fmt.Sprintf("ls %s | peco", txbufDir)
	c := exec.Command("bash", "-c", cmd)
	o, err := c.Output()
	if err != nil {
		return "", err
	}
	return string(o), nil
}
