package goshell

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/skeptycal/gosimple/datatools/bufferpool"
)

const (
	defaultShell = "bash"
	cOption      = "-c"
)

var (
	shell        string          = Getenv("$SHELL", defaultShell)
	argStubBlank []string        = []string{cOption}
	ctxShell     context.Context = context.TODO()
)

// func getCmd()

// func RunCMD(name string, args ...string) (err error, stdout, stderr []string) {
// 	c := cmd.NewCmd(name, args...)
// 	s := <-c.Start()
// 	stdout = s.Stdout
// 	stderr = s.Stderr
// 	return
// }
func bufferShell(ctx context.Context, stdout, stderr *bytes.Buffer, name string, args ...string) (string, string, error) {
	cmd := exec.CommandContext(ctxShell, name, argStub(args...)...) //append(argStubBlank, args...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

// ShBuffered executes the named shell command and returns
// stdout, stderr, and any error that occurred.
//
// The buffers for the command outputs are standard
// bytes.Buffers and may not be performant for multiple
// repeated commands. See ShellPool.
//
// The default timeout is 10 seconds.
func Shell(commands ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	// cmd := exec.Command(shell, "-c", commands)
	// cmd := exec.CommandContext(ctxShell, shell, append(argStubBlank, commands...)...)
	return bufferShell(ctxShell, &stdout, &stderr, commands[0], commands[1:]...)

	// cmd.Stdout = &stdout
	// cmd.Stderr = &stderr
	// err := cmd.Run()
	// return stdout.String(), stderr.String(), err
}

var buf = bufferpool.NewPool[*bytes.Buffer]()

var syncPool = sync.Pool{}

// ShellPool executes the named shell command and returns
// stdout, stderr, and any error that occurred.
//
// The buffers for the command outputs are sync.Pool
// bytes.Buffers and are safe for concurrent use and
// streamlined for high performance.
//
// The default timeout is 10 seconds.
func ShellPool(args ...string) (string, string, error) {
	var stdout *bytes.Buffer
	var stderr *bytes.Buffer
	defer bufferpool.Swimmer(syncPool, stdout)()
	defer bufferpool.Swimmer(syncPool, stderr)()
	//  = bufferPool.Get().(bytes.Buffer)
	// defer bufferPool.Put(stdout)

	// cmd := exec.Command("sh", "-c", `dscl -q . -read /Users/"$(whoami)" NFSHomeDirectory | sed 's/^[^ ]*: //'`)
	return bufferShell(ctxShell, stdout, stderr, shell, argStub(args...)...)

	cmd := exec.CommandContext(ctxShell, shell, append(argStubBlank, args...)...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	err := cmd.Run()
	sout := strings.TrimSpace(stdout.String())
	serr := strings.TrimSpace(stderr.String())
	// pid := cmd.Process.Pid

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
		return sout, serr, fmt.Errorf("error executing shell command (pid: %d): %v", cmd.Process.Pid, err)
	}
	if serr != "" {
		return sout, serr, fmt.Errorf("stderr (pid: %d): %20v", cmd.Process.Pid, err)
	}
	return sout, "", nil
}

func argStub(args ...string) []string {
	return append(argStubBlank, args...)
}
