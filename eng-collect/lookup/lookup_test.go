package lookup

import "testing"

func TestLookUpWord(t *testing.T) {
	testcases := map[string]struct {
		word          string
		expectedWord  Word
		expectedError error
	}{
		"case-1: word exists": {
			word: "cookie",
			expectedWord: Word{
				Name:       "cookie",
				Level:      "A2",
				Definition: "a small, flat, sweet food made from flour and sugar",
				Example:    "How clever of you to buy chocolate chip cookies - they're my favourites.",
			},
			expectedError: nil,
		},
		"case-2: word exists": {
			word: "milk",
			expectedWord: Word{
				Name:       "milk",
				Level:      "A1",
				Definition: "the white liquid produced by cows, goats, and sheep and used by humans as a drink or for making butter, cheese, etc",
				Example:    "Do you take milk in your tea?",
			},
			expectedError: nil,
		},
	}

	for _, tc := range testcases {
		result, err := LookUpWord(tc.word)
		if err != nil {
			t.Errorf("error when look up '%s' should be nil, error: %v\n", tc.word, err)
			return
		}

		if result.Definition != tc.expectedWord.Definition {
			t.Errorf("expected definition: %s, got: %s", tc.expectedWord.Definition, result.Definition)
		}

		if result.Level != tc.expectedWord.Level {
			t.Errorf("expected level: %s, got: %s", tc.expectedWord.Level, result.Level)
		}

		if result.Example != tc.expectedWord.Example {
			t.Errorf("expected example: %s, got: %s", tc.expectedWord.Example, result.Example)
		}
	}
}
