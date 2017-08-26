package cmds

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestCmds(t *testing.T) {
	cmd := exec.Command("go", "build")
	cmd.Stdin = strings.NewReader("build")

	var out = bytes.NewBuffer(nil)
	var eout = bytes.NewBuffer(nil)

	cmd.Stdout = out
	cmd.Stderr = eout

	err := cmd.Run()
	fmt.Println("out:", out.String())
	fmt.Println("-----------1-----------", err)
	fmt.Println("eout:", eout.String())

	cmd.Stdin = strings.NewReader("build")
	cmd.Run()
	fmt.Println("----------2--------------", err)
	fmt.Println("eout:", eout.String())

}
