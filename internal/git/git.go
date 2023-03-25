package git

import (
	"fmt"
	"os/exec"
)

type GitError struct {
	msg string
	err error
}

func (e *GitError) Error() string {
	return fmt.Sprintf("Error from git: %s (%s)", e.msg, e.err)
}
func (e *GitError) Unwrap() error {
	return e.err
}

func RunGit(txbufDir string, args []string) (string, error) {
	cmd := fmt.Sprintf("cd %s && git", txbufDir)
	for _, arg := range args {
		cmd = cmd + " " + arg
	}
	c := exec.Command("bash", "-c", cmd)
	o, err := c.Output()
	if err != nil {
		return "", &GitError{err: err, msg: fmt.Sprintf("`%s`の実行に失敗", cmd)}
	}
	return string(o), nil
}
