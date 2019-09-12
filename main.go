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
	fmt.Printf("Running %v as PID:%d\n", os.Args[2:], os.Getpid())
	// cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

	// redirect all in/out file descript to std
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdin

	// setup namespaces
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// CLONE_NEWUTS new hostname
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("Running %v as PID:%d\n", os.Args[2:], os.Getpid())
	must(syscall.Sethostname([]byte("newhost")))
	must(syscall.Chroot("./chroot"))
	must(syscall.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	defer syscall.Unmount("proc", 0)
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
