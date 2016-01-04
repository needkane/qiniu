package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func run() *exec.Cmd {
	cmd := exec.Command("go", "run", "childProcess.go")
	cmd.Start()
	return cmd
}

func main() {
	cmd := run()
	//	var i int
	//	fmt.Scanf("%d",&i)
	time.Sleep(time.Second * 1)
	fmt.Println("1---------", cmd.Process.Pid)
	pid := cmd.Process.Pid
	syscall.ki
	proc, _ := os.FindProcess(cmd.Process.Pid)
	proc.Signal(syscall.Signal(0))
	cmd.Process.Wait()

	fmt.Println("2---------", cmd.Process.Pid)
	time.Sleep(time.Second * 16)

	// Child process (server.go) is still running!

	fmt.Println("3---------", cmd.Process.Pid)
}
