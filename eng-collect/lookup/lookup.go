package lookup

import (
	"errors"
	"strings"

	"github.com/gocolly/colly/v2"
)

const CrawlURL = "https://dictionary.cambridge.org/dictionary/english/"

func LookUpWord(word string) (*Word, error) {
	c := colly.NewCollector()
	w := Word{Name: word}

	// Google Chrome on Windows
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36"

	// get definition
	firstDefinitionFound := false
	c.OnHTML(".def.ddef_d.db", func(e *colly.HTMLElement) {
		if !firstDefinitionFound {
			firstDefinitionFound = true
			text := strings.TrimSpace(e.Text)
			if text[len(text)-1] == ':' {
				w.Definition = text[:len(text)-1]
			} else {
				w.Definition = text
			}
		}
	})

	// get level (if it exists)
	c.OnHTML(".epp-xref.dxref", func(e *colly.HTMLElement) {
		if e.Index == 0 {
			w.Level = e.Text
		}
	})

	// get example (if it exists)
	c.OnHTML(".eg.dexamp.hax", func(e *colly.HTMLElement) {
		if e.Index == 0 {
			w.Example = e.Text
		}
	})

	var err error
	c.OnScraped(func(r *colly.Response) {
		if w.Definition == "" {
			err = ErrWordNotFound
			return
		}
	})
	url := CrawlURL + word

	c.Visit(url) // by using colly, you only start web scraping after building callback functions to process crawl data

	if !firstDefinitionFound {
		err = ErrDenyCrawl
	}

	if err != nil {
		return nil, err
	}

	return &w, nil
}

type Word struct {
	Name       string
	Level      string
	Definition string
	Example    string
}

var ErrWordNotFound = errors.New("word not found")
var ErrDenyCrawl = errors.New("being denied to crawl")
