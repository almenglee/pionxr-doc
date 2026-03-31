//go:build debug

package main

func unimplementedFeature(msg string) error {
	panic("unimplemented error: " + msg)
}

func undecidedError(msg string) error {
	panic("undecided error: " + msg)
}

func bug(msg string) error {
	panic("bug: " + msg)
}

func warn(msg string) {
	println("warning: " + msg)
}
