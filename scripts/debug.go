//go:build debug

package main

func unimplementedError(msg string) error {
	panic("unimplemented error: " + msg)
}

func fatal(msg string) error {
	panic("fatal program error: " + msg)
}

func warn(msg string) {
	println("warning: " + msg)
}
