package main

import (
	"fmt"
	"io"
	"os"
	"pionxr-doc/internal/builder"
)

func main() {
	if err := Run(os.Stdin, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Run(r io.Reader, w io.Writer, errOut io.Writer) error {
	src, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	result := builder.NewBuilder(string(src)).Build()

	err = result.EncodeJSON(w)
	if err != nil {
		return err
	}

	return nil
}
