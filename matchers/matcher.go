package matchers

import (
	"fmt"
)

type Matcher interface {
	Example() (input string, description string)
	Matches(input string) bool
	Run(input string) (finalValue int, computation string, err error)
}

func All() []Matcher {
	return []Matcher{
		newSimpleMatcher(),
	}
}

func AllExamples() string {
	matchers := All()
	ret := ""
	for _, m := range matchers {
		input, desc := m.Example()
		ret += fmt.Sprintf("\t%s: %s\n", input, desc)
	}
	return ret
}
