package pirate

import (
	"fmt"
	"net/url"

	"github.com/gocolly/colly"
)

type LeetxSearch struct{}

func (*LeetxSearch) Search(search string) ([]Metadata, error) {

	c := colly.NewCollector()

	var metadata []Metadata

	c.OnHTML("table tr td:nth-of-type(1) a:nth-of-type(2)", func(e *colly.HTMLElement) {
		metadata = append(metadata, Metadata{
			Name: e.Text,
			Hash: e.Attr("href"),
		})
	})

	c.Visit(fmt.Sprintf("https://www.1337xx.to/search/%s/1/", url.QueryEscape(search)))

	return metadata, nil
}

func (*LeetxSearch) GetMagnet(metadata Metadata) string {

	c := colly.NewCollector()
	var result string

	c.OnHTML("#down_magnet", func(e *colly.HTMLElement) {
		result = e.Attr("href")
	})

	c.Visit(fmt.Sprintf("https://www.1337xx.to%s", metadata.Hash))

	return result
}

func (*LeetxSearch) GetName() string {
	return "1337x"
}
