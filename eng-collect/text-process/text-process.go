package textprocess

import (
	"log/slog"
	"strings"

	"github.com/aaaton/golem/v4"
	"github.com/aaaton/golem/v4/dicts/en"
	"github.com/bbalet/stopwords"
)

func Tokenize(text string) []string {

	text = Lemmatize(text)
	text = RemoveStopword(text)

	rawTokens := strings.Split(text, " ")

	tokens := []string{}

	for _, token := range rawTokens {
		if token != "" {
			tokens = append(tokens, token)
		}
	}

	return tokens
}

func RemoveStopword(text string) string {
	return stopwords.CleanString(text, "en", true)
}

func Lemmatize(text string) string {
	lemmatizer, err := golem.New(en.New())
	if err != nil {
		slog.Info("fail to create lemmatizer", "error", err)

		// if lemmatize not working just ignore, let cambridge dictionary do that ... (sorry if you have read this)
		return text
	}
	return lemmatizer.Lemma(text)
}
