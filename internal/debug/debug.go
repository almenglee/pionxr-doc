//go:build debug

package debug

func UnimplementedFeature(msg string) error {
	panic("unimplemented error: " + msg)
}

func UndecidedError(msg string) error {
	panic("undecided error: " + msg)
}

func Bug(msg string) error {
	panic("bug: " + msg)
}

func Warn(msg string) {
	println("warning: " + msg)
}
