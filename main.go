package main

import (
	"fmt"
	"os"
)

const (
	fmtUsage = "Usage: %s filepath start:stop:out...\n"
)

func main() {
	os.Exit(run())
}

func run() (exitCode int) {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, fmtUsage, os.Args[0])
		exitCode = 1
		return
	}

	filePath := os.Args[1]
	slices := os.Args[2:]

	slicer, err := NewSlicer(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		exitCode = 2
		return
	}

	for number, slice := range slices {
		err = slicer.Do(number, slice)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			exitCode = 3
			break
		}
	}

	err = slicer.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if exitCode == 0 {
			exitCode = 4
		}
	}

	return
}
