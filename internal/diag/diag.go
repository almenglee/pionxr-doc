package diag

import (
	"fmt"
)

type Severity string

const (
	Info    Severity = "info"
	Warning Severity = "warning"
	Error   Severity = "error"
)

type Diagnostic struct {
	Severity Severity
	Message  string
	Pos      Position
}

type Position interface {
	Pos() string
	String() string
}

func (d Diagnostic) Error() string {
	if d.Pos == nil {
		return d.Message
	}
	return fmt.Sprintf("%s: %s", d.Pos.String(), d.Message)
}

func Errorf(format string, args ...interface{}) *Diagnostic {
	return &Diagnostic{
		Severity: Error,
		Message:  fmt.Sprintf(format, args...),
	}
}

type Collector struct {
	list   []*Diagnostic
	lut    map[*Diagnostic]struct{}
	hasErr bool
}

func NewCollector() *Collector {
	return &Collector{
		list: []*Diagnostic{},
		lut:  make(map[*Diagnostic]struct{}),
	}
}

func (c *Collector) HasErrors() bool {
	return c.hasErr
}

func (c *Collector) Add(d *Diagnostic) {
	if d.Severity == Error {
		c.hasErr = true
	}
	if _, exists := c.lut[d]; !exists {
		c.list = append(c.list, d)
		c.lut[d] = struct{}{}
	}
}

func (c *Collector) List() []*Diagnostic {
	out := make([]*Diagnostic, len(c.list))
	copy(out, c.list)
	return out
}
