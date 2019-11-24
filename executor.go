package ago

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// Executor Interface for run command.
type Executor interface {
	Exec(name string, args []string) error
}

// DefaultExecutor Default executor with io.Writer to stdout for commands run.
type DefaultExecutor struct {
	Writer io.Writer
}

func (e *DefaultExecutor) Exec(name string, args []string) error {
	stderr := &bytes.Buffer{}
	if e.Writer == nil {
		e.Writer = os.Stdout
	}
	cmd := exec.Command(name, args...)
	cmd.Stderr = stderr
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Fprintf(e.Writer, "%s\n", scanner.Text())
		}
	}()
	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
}
