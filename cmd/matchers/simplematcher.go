package matchers

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/sirupsen/logrus"
)

var pattern = regexp.MustCompile(`^([0-9]*)d([0-9]*)([+-]?[0-9]*)$`)

type simpleMatcher struct{}

func newSimpleMatcher() Matcher {
	return simpleMatcher{}
}

func (sm simpleMatcher) Example() (string, string) {
	return "1d20+4", "(1) 20-sided die, with a +4 modifier"
}

func (sm simpleMatcher) Matches(input string) bool {
	return pattern.MatchString(input)
}

func (sm simpleMatcher) Run(input string) (int, string, error) {
	// parse the input
	parts := pattern.FindStringSubmatch(input)
	logrus.Debug(parts)
	numDice, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, "", err
	}
	dieMax, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, "", err
	}
	// plusMinus := parts[3]
	modifier := 0
	if len(parts[3]) > 0 {
		modifier, err = strconv.Atoi(parts[3])
		if err != nil {
			return 0, "", err
		}
	}

	// compute!
	sum := 0
	calca := ""
	for i := 0; i < numDice; i++ {
		roll := r(dieMax)
		if i > 0 {
			calca += "+"
		}
		calca += fmt.Sprintf("(%d)", roll)
		sum += roll
	}

	// add modifier
	if modifier < 0 {
		calca += fmt.Sprintf("%d", modifier)
	} else if modifier > 0 {
		calca += fmt.Sprintf("+%d", modifier)
	}
	sum += modifier

	// pretty it up
	calca = fmt.Sprintf("%d = %s", sum, calca)

	return sum, calca, nil
}
