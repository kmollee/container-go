package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
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
		// CLONE_NEWUTS: new hostname
		// CLONE_NEWNS: Unshare the mount namespace, so that the calling process has
		// a private copy of its namespace which is not shared with any other process
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS, // not visible to host
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("Running %v as PID:%d\n", os.Args[2:], os.Getpid())

	cg("newhost")
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

// create cgroup
func cg(name string) {
	pids := "/sys/fs/cgroup/pids"
	os.Mkdir(filepath.Join(pids, name), 0755)

	must(ioutil.WriteFile(filepath.Join(pids, name, "pids.max"), []byte("20"), 0700))
	must(ioutil.WriteFile(filepath.Join(pids, name, "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
