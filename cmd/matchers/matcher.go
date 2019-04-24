package matchers

import (
	"fmt"
	"math/rand"
	"time"
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

func r(max int) int {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	// instead of returning [0,max), we want to return [1,max]
	return rand.Intn(max) + 1
}
