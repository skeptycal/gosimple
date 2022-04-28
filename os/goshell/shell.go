package goshell

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const defaultShell = "bash"

var shell = GetEnv("$SHELL", defaultShell)

func GetEnv(envVarName string, defaultValue string) (retval string) {
	retval = os.ExpandEnv(envVarName)
	if retval == "" {
		return defaultValue
	}
	return
}

// func getCmd()

// func RunCMD(name string, args ...string) (err error, stdout, stderr []string) {
// 	c := cmd.NewCmd(name, args...)
// 	s := <-c.Start()
// 	stdout = s.Stdout
// 	stderr = s.Stderr
// 	return
// }

func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(shell, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

const cOption = "-c"

var argStub = []string{cOption}

var ctxShell = context.TODO()

// ShBuffered executes the named shell command and returns
// stdout, stderr, and any error that occurred.
//
// The default timeout is 10 seconds.
func ShBuffered(args ...string) (string, string, error) {
	var stdout *bytes.Buffer
	var stderr *bytes.Buffer
	defer Swimmer(stdout)(stdout)
	defer Swimmer(stderr)(stderr)
	//  = bufferPool.Get().(bytes.Buffer)
	// defer bufferPool.Put(stdout)

	// cmd := exec.Command("sh", "-c", `dscl -q . -read /Users/"$(whoami)" NFSHomeDirectory | sed 's/^[^ ]*: //'`)
	cmd := exec.CommandContext(ctxShell, shell, append(argStub, args...)...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	err := cmd.Run()
	sout := strings.TrimSpace(stdout.String())
	serr := strings.TrimSpace(stderr.String())
	pid := cmd.Process.Pid

	/* data (cmd.ProcessState) methods:
	String
	Stdout
	Pid
	ExitCode
	Exited
	Success
	Sys
	SysUsage
	SystemTime
	UserTime
	SystemTime().Hours
	SystemTime().Microseconds
	SystemTime().Milliseconds
	*/

	// TODO process data from shell commands
	// data := cmd.ProcessState
	// stime := data.SystemTime()
	// utime := data.UserTime()

	if err != nil {
		return sout, serr, fmt.Errorf("error executing shell command (pid: %d): %v", pid, err)
	}
	if serr != "" {
		return sout, serr, fmt.Errorf("stderr reported error (pid: %d): %20v", pid, err)
	}
	return sout, "", nil
}
