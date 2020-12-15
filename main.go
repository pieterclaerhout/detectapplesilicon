package main

import (
	"fmt"
	"syscall"
)

func main() {

	r, err := syscall.Sysctl("sysctl.proc_translated")
	if err.Error() == "no such file or directory" {
		fmt.Println("Running on intel mac")
	}

	if r == "\x00\x00\x00" {
		fmt.Println("Running on apple silicon natively")
	}

	if r == "\x01\x00\x00" {
		fmt.Println("Running on apple silicon under Rosetta 2")
	}

}
