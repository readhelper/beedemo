package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
	"os"
	"github.com/readhelper/beedemo/logger"
)

func TestRunDefault(t *testing.T) {
	// don't launch etcdctl when invoked via go test
	if strings.HasSuffix(os.Args[0], "etcdctl.test") {
		logger.Info("don't launch etcdctl when invoked via go test")
		return
	}
	main()
}

func TestVersion(t *testing.T) {
	runApp([]string{"cmds", "-v"})
	//var out = bytes.NewBuffer(nil)
	//logger.Info("out:", out.String())
}

func TestCmds(t *testing.T) {
	cmd := exec.Command("go", "build")
	cmd.Stdin = strings.NewReader("build")

	var out = bytes.NewBuffer(nil)
	var eout = bytes.NewBuffer(nil)

	cmd.Stdout = out
	cmd.Stderr = eout

	err := cmd.Run()
	logger.Info("out:", out.String(), ",eout:", eout.String(), "err:", err)
}
