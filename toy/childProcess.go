package main

import (
	"fmt"
	"os/exec"
	"time"
)

func run() *exec.Cmd {
	cmd := exec.Command("go", "run", "server2.go")
	cmd.Start()
	return cmd
}

func main() {
	cmd := run()
	fmt.Println("child---------", cmd.Process.Pid)

	time.Sleep(time.Second * 40)
	cmd.Process.Wait()
	fmt.Println("child------------", cmd.Process.Pid)
}
