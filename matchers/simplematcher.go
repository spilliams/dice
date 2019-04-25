package matchers

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spilliams/dice/dice"
)

var simplePattern = regexp.MustCompile(`^([0-9]+)d([0-9]+)([+-]{1}[0-9]+)?$`)

type simpleMatcher struct{}

func newSimpleMatcher() Matcher {
	return simpleMatcher{}
}

func (sm simpleMatcher) Example() (string, string) {
	return "1d20+4", "(1) 20-sided die, with a +4 modifier"
}

func (sm simpleMatcher) Matches(input string) bool {
	return simplePattern.MatchString(input)
}

func (sm simpleMatcher) Run(input string) (int, string, error) {
	// parse the input
	parts := simplePattern.FindStringSubmatch(input)
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
	rolls := dice.NdM(numDice, dieMax)
	sum := 0
	calca := ""
	for i, roll := range rolls {
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
