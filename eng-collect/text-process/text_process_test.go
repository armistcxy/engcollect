package textprocess

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	testcases := map[string]struct {
		input    string
		expected []string
	}{
		"case1": {
			input:    "A base Technology for the future is being created",
			expected: []string{"base", "technology", "future", "created"},
		},
	}

	for testName, tc := range testcases {
		t.Run(testName, func(t *testing.T) {
			result := Tokenize(tc.input)
			for _, token := range result {
				t.Logf("'%s'", token)
			}
			if len(result) != len(tc.expected) {
				t.Errorf("Length of tokens array should be: %v, got: %v", len(tc.expected), len(result))
				return
			}
			for i := range result {
				if result[i] != tc.expected[i] {
					t.Errorf("Expected token: %s at position %v, got: %s", tc.expected[i], i, result[i])
				}
			}
		})
	}
}
