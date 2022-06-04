package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/skeptycal/gosimple/cli"
	"github.com/skeptycal/gosimple/cli/terminal"
)

func main() {
	w := terminal.GetWinSize()
	if w == nil {
		log.Fatal("winsize not available")
	}

	fmt.Println("GetWinSize(): ", w)
	fmt.Println("CheckIfTerminal(): ", cli.CheckIfTerminal(os.Stdout))
	fmt.Println("Cols(): ", cli.Cols())
	fmt.Println("Rows(): ", cli.Rows())
	fmt.Println("XPixels(): ", cli.XPixels())
	fmt.Println("YPixels(): ", cli.YPixels())

	env("$PATH")
	env("$SHELL")
	shell("date")
	shell("echo $PATH")
	echo("$PATH")
	fmt.Println(os.ExpandEnv("PATH is ${PATH}."))
}

func env(s string) {
	v := os.ExpandEnv(s)
	if v == "" {
		v = execEnv(s)
	}
	fmt.Printf("%q: %s\n", s, v)
}

func execEnv(s string) string {
	b, _ := exec.Command("echo", s).Output()
	retval := cli.B2S(b)
	fmt.Printf("%s: %s", s, retval)
	return retval
}

func echo(args ...string) (string, error) {
	commandstring := fmt.Sprintf("$(echo %s)", strings.Join(args, " "))
	fmt.Println("commandstring: ", commandstring)
	out, err := exec.Command("eval", commandstring).Output()
	if err != nil {
		return "", err
	}
	fmt.Printf("echo %s: %s\n", strings.Join(args, " "), out)

	return "", nil
}

func shell(args ...string) (string, error) {
	out, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		return "", err
	}
	fmt.Printf("%s: %s\n", strings.Join(args, " "), out)
	return "", nil
}
