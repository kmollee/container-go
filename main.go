package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// docker run <image> <cmd> <params>
// go run main.go run <cmd> <params>

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("Not implment")
	}
}

func run() {
	fmt.Printf("Running %v\n", os.Args[2:])
	// cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

	// redirect all in/out file descript to std
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdin

	// setup namespaces
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// CLONE_NEWUTS new hostname
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	must(cmd.Run())
}

func child() {

	fmt.Println("child")
	must(syscall.Sethostname([]byte("newhost")))
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	// redirect all in/out file descript to std
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdin

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
