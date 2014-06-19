package main

import (
	"fmt"
	"os/exec"
)

type Runner interface {
	Run(string, ...string) ([]byte, error)
}

type RealRunner struct{}

var runner Runner

func (r RealRunner) Run(command string, args ...string) ([]byte, error) {
	out, err := exec.Command(command, args...).CombinedOutput()
	return out, err
}

func Hello() string {
	out, err := runner.Run("echo", "hello")
	if err != nil {
		panic(err)
	}
	return string(out)
}

func main() {
	runner = RealRunner{}
	fmt.Println(Hello())
}
