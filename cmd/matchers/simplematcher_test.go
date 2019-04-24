package matchers

import (
	"fmt"
	"strings"
	"testing"
)

func TestSimpleMatcher(t *testing.T) {
	cases := []struct {
		input         string
		expectMatches bool
		expectError   string
	}{
		{
			input:         "1d20",
			expectMatches: true,
		},
		{
			input:         "3d12",
			expectMatches: true,
		},
		{
			input:         "1d10+4",
			expectMatches: true,
		},
		{
			input:         "1d10-199",
			expectMatches: true,
		},
	}

	sm := newSimpleMatcher()

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			if sm.Matches(c.input) != c.expectMatches {
				if c.expectMatches {
					t.Fatalf("simplematcher should match %s, but doesn't", c.input)
				}
				t.Fatalf("simplematcher should not match %s, but does", c.input)
			}
			_, _, err := sm.Run(c.input)
			err = checkError(err, c.expectError)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func checkError(actual error, expected string) error {
	if actual == nil {
		if len(expected) > 0 {
			return fmt.Errorf("expected error %s, but got none", expected)
		}
		// actual is nil, expected is empty. ok
		return nil
	}
	// actual is not nil
	if len(expected) == 0 {
		return fmt.Errorf("actual error is %v, but we expected no error", actual)
	}
	actualTrimmed := strings.TrimSpace(actual.Error())
	expectedTrimmed := strings.TrimSpace(expected)
	if actualTrimmed != expectedTrimmed {
		return fmt.Errorf("actual error is %s, but we expected %s", actualTrimmed, expectedTrimmed)
	}
	return nil
}
