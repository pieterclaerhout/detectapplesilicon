package main

import (
	"fmt"
	"runtime"
	"syscall"
)

func main() {

	r, err := syscall.Sysctl("sysctl.proc_translated")
	if err != nil && err.Error() == "no such file or directory" {
		fmt.Println("Running on intel mac, arch:", runtime.GOARCH)
	}

	if r == "\x00\x00\x00" {
		fmt.Println("Running on apple silicon natively, arch:", runtime.GOARCH)
	}

	if r == "\x01\x00\x00" {
		fmt.Println("Running on apple silicon under Rosetta 2, arch:", runtime.GOARCH)
	}

}
