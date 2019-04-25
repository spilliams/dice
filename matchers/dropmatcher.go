package matchers

import "regexp"

var dropPattern = regexp.MustCompile(`^([0-9]+)d([0-9]+) drop (lo[^ ]*|hi[^ *])( [0-9]+)?$`)

type dropMatcher struct{}

func newDropMatcher() Matcher {
	return dropMatcher{}
}

func (dm dropMatcher) Example() (string, string) {
	return "4d6 drop lowest 3", "A roll term \"drop\" lowest (or highest) with an optional count for how many to drop"
}

func (dm dropMatcher) Matches(input string) bool {
	return dropPattern.MatchString(input)
}

func (dm dropMatcher) Run(intput string) (int, string, error) {
	return 0, "", nil
}
