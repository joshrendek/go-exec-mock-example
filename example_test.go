package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

type TestRunner struct{}

func (r TestRunner) Run(command string, args ...string) ([]byte, error) {
	cs := []string{"-test.run=TestHelperProcess", "--"}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	out, err := cmd.CombinedOutput()
	return out, err
}

func TestHello(t *testing.T) {
	runner = TestRunner{}
	out := Hello()
	if out == "testing helper process" {
		t.Logf("out was eq to %s", string(out))
	}
}

func TestHelperProcess(*testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	defer os.Exit(0)
	fmt.Println("testing helper process")
}
