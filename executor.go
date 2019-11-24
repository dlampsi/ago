package ago

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
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
	exitCode, err := e.ansibleExec(name, args)
	switch exitCode {
	case 0: // OK or no hosts matched
		return nil
	case 1: // Error
		return err
	case 2: // One or more hosts failed
		return fmt.Errorf("one or more hosts failed: %v", err)
	case 3: // One or more hosts were unreachable
		return fmt.Errorf("one or more hosts were unreachable: %v", err)
	case 4: // Parser error
		return err
	case 5: // Bad or incomplete options
		return fmt.Errorf("bad or incomplete options %v", err)
	case 99: // User interrupted execution
		return fmt.Errorf("user interrupted execution: %v", err)
	case 250: // Unexpected error
		return fmt.Errorf("unexpected error: %v", err)
	default: // Unknown error
		return fmt.Errorf("unknown error: %v", err)
	}
}

// Execute ansible command. Return exit code and error.
// Process all error codes.
func (e *DefaultExecutor) ansibleExec(name string, args []string) (int, error) {
	if e.Writer == nil {
		e.Writer = os.Stdout
	}

	// Fake exit code for return in direct error cases
	exitCode := 555

	cmd := exec.Command(name, args...)

	// Scan Stderr
	cmdStderr, err := cmd.StderrPipe()
	if err != nil {
		return exitCode, err
	}
	errrScan := bufio.NewScanner(cmdStderr)
	go func() {
		for errrScan.Scan() {
			fmt.Fprintf(e.Writer, "%s\n", errrScan.Text())
		}
	}()

	// Scan Stdout
	cmdStdout, err := cmd.StdoutPipe()
	if err != nil {
		return exitCode, err
	}
	outScan := bufio.NewScanner(cmdStdout)
	go func() {
		for outScan.Scan() {
			fmt.Fprintf(e.Writer, "%s\n", outScan.Text())
		}
	}()

	// Start command
	err = cmd.Start()
	if err != nil {
		return exitCode, err
	}

	// Wait for command finish
	err = cmd.Wait()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if stat, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				exitCode = stat.ExitStatus()
			}
		}
		return exitCode, err
	}

	waitStat := cmd.ProcessState.Sys().(syscall.WaitStatus)
	exitCode = waitStat.ExitStatus()

	return exitCode, nil
}
