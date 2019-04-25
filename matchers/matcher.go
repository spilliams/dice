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
		newLongMatcher(),
		newDropMatcher(),
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

/*
package matchers

import "regexp"

var fooPattern = regexp.MustCompile(`^$`)

type fooMatcher struct{}

func newFooMatcher() Matcher {
	return fooMatcher{}
}

func (fm fooMatcher) Example() (string, string) {
	return "1d4", "(1) 4-sided die"
}

func (fm fooMatcher) Matches(input string) bool {
	return fooPattern.MatchString(input)
}

func (fm fooMatcher) Run(input string) (int, string, error) {
	return 0, "", nil
}
*/
