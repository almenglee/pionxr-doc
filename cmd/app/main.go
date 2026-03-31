package main

import (
	"os"
	"os/exec"
)

func parseDocument(path string) []byte {
	exec_path := "./core"

	src, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(exec_path)
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	go func() {
		defer stdin.Close()
		_, _ = stdin.Write(src)
	}()

	result, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	return result

}

func main() {
	result := parseDocument("../draft.md")
	os.Stdout.Write(result)

}
