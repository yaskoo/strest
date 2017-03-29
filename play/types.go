package play

import (
	"fmt"
)

const (
	Status ResProp = iota
	Header
	Body
)

type ResProp int

type DiffEntry struct {
	Prop     ResProp
	Expected interface{}
	Actual   interface{}
}

func (r ResProp) String() string {
	switch r {
	case Status:
		return "Status codes"
	case Header:
		return "Headers"
	case Body:
		return "Bodies"
	default:
		return "<unknown>"
	}
}

func (d DiffEntry) String() string {
	return fmt.Sprintf("Expected %v: %v, but got: %v", d.Prop, d.Expected, d.Actual)
}

type Response struct {
	Status  int                 `yaml:"status"`
	Headers map[string][]string `yaml:"headers"`
	Body    string              `yaml:"body"`
}

type Context struct {
	step      int
	Responses []*Response
}

type Step struct {
	Name     string              `yaml:"name"`
	Url      string              `yaml:"url"`
	Method   string              `yaml:"method"`
	Headers  map[string][]string `yaml:"headers"`
	Body     string              `yaml:"body"` // TODO: make this []byte
	Response *Response           `yaml:"response"`
	Expect   *Response           `yaml:"expect"`
}

type Play struct {
	testMode bool
	Steps    []Step `yaml:"steps"`
}
