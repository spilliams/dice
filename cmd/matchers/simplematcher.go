package matchers

import "fmt"

type simpleMatcher struct {
}

func newSimpleMatcher() Matcher {
	return simpleMatcher{}
}

func (sm simpleMatcher) Example() (string, string) {
	return "1d20", "(1) 20-sided die"
}

func (sm simpleMatcher) Matches(input string) bool {
	return input == "1d20"
}

func (sm simpleMatcher) Run(input string) (int, string, error) {
	n := r(20)
	return n, fmt.Sprintf("%v", n), nil
}
