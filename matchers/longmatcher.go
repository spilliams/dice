package matchers

import "regexp"

var longPattern = regexp.MustCompile(`^$`)

type longMatcher struct{}

func newLongMatcher() Matcher {
	return longMatcher{}
}

func (lm longMatcher) Example() (string, string) {
	return "1d4+5d6-3+...", "Any number of roll and modifier terms, separated by a + or -"
}

func (lm longMatcher) Matches(input string) bool {
	return longPattern.MatchString(input)
}

func (lm longMatcher) Run(input string) (int, string, error) {
	return 0, "", nil
}
