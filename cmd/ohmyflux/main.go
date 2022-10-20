package main

import "os"

type exitCode int

const (
	exitOk exitCode = 0
)

func main() {
	code := mainRun()
	os.Exit(int(code))
}

func mainRun() exitCode {
	return exitOk
}
