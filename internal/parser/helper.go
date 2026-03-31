package parser

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func readDraftLines() ([]string, string, error) {
	paths := []string{"draft.md", "../draft.md"}

	var file *os.File
	var err error

	for _, path := range paths {
		file, err = os.Open(path)
		if err == nil {
			defer file.Close()
			lines, scanErr := scanLines(file)
			if scanErr != nil {
				return nil, "", scanErr
			}
			return lines, path, nil
		}
	}

	return nil, "", fmt.Errorf("could not open draft.md from known paths: %v", paths)
}

func scanLines(file *os.File) ([]string, error) {
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func joinBlockRawText(lines []string, start, end int) string {
	if start >= end {
		return ""
	}

	return strings.Join(lines[start:end], "\n")
}

func syntaxError(format string, args ...any) error {
	return fmt.Errorf(format, args...)
}

func joinErrors(errs []error) error {
	if len(errs) == 0 {
		return nil
	}

	return errors.Join(errs...)
}

func isLetter(ch rune) bool  { return 'a' <= lower(ch) && lower(ch) <= 'z' || ch == '_' }
func lower(ch rune) rune     { return ('a' - 'A') | ch }
func isDecimal(ch rune) bool { return '0' <= ch && ch <= '9' }
func isIdent(str string) bool {
	if str == "" {
		return false
	}
	if !isLetter(rune(str[0])) {
		return false
	}

	for _, ch := range str {
		if !isLetter(rune(ch)) && !isDecimal(rune(ch)) {
			return false
		}
	}

	return true
}

func ident(str string) string {
	for i, ch := range str {
		if !isLetter(rune(ch)) && !isDecimal(rune(ch)) {
			return str[:i]
		}
	}

	return str
}
