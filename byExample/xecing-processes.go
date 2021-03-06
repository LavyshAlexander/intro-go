package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	path := os.Getenv("PATH")
	fmt.Println(path)

	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()
	fmt.Println(env)

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
