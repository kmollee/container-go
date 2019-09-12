package main

import (
	"os"
	"os/exec"
)

// docker run <image> <cmd> <params>
// go run main.go run <cmd> <params>

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("Not implment")
	}
}

func run() {
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
