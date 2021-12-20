package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	// "net/http"
	// "net/http/cookiejar"
	// "net/url"
	// "github.com/PuerkitoBio/goquery"
)

var Parsers = map[string]func(string) PageOutput{
	"ParseSANS":       ParseSANS,
	"ParseTrendMicro": ParseTrendMicro,
}

// Article returns the output of parsing each article on the webpage.
type Article struct {
	title string
	link  string
	desc  string
	date  string
}

type PageOutput struct {
	parser   string
	articles []Article
}

func fetchPage(url string) (string, error) {
	// Make HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	// Convert the output to string
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("unable to read the http response body into io: %v", resp)
	}

	return string(b), nil
}

// ParseTrendMicro visits the TrendMicro research site and returns the article meta data presented therein.
func ParseTrendMicro(url string) PageOutput {
	contextLogger.Info(fmt.Sprintf("Processing %s using ParseTrendMicro", url))

	resp, err := fetchPage(url)
	if err != nil {
		contextLogger.Warnf(fmt.Sprintf("ParseTrendMicro: Error occured fetching url(%s): %v", url, err))
	}

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(resp))
	if err != nil {
		contextLogger.Warnf(fmt.Sprintf("ParseTrendMicro: Error occured fetching url(%s): %v", url, err))
	}

	// Aggregate all the titles gotten from this page.
	var articles []Article

	// Traverse the article text containers and extract what is needed.
	dom.Find("div.text-container").Each(func(index int, element *goquery.Selection) {
		var out Article

		element.Find(".heading").Each(func(index int, element *goquery.Selection) {
			out.title = element.Text()
		})
		element.Find("a[href$='.html']").Each(func(index int, element *goquery.Selection) {
			out.link, _ = element.Attr("href")
		})
		element.Find("p.description").Each(func(index int, element *goquery.Selection) {
			out.desc = strings.TrimSpace(element.Text())
		})
		// Get the date from eth span element. TODO(igbe): Update below for a better way to pull date.
		element.Find("span:contains(', 20')").Each(func(index int, element *goquery.Selection) {
			out.date = element.Text()
		})

		articles = append(articles, out)
	})

	return PageOutput{parser: "ParseTrendMicro", articles: articles}
}

// TODO(igbe): Revisit this.
func ParseSANS(url string) PageOutput {
	contextLogger.Info(fmt.Sprintf("Processing %s using ParseSANS", url))
	return PageOutput{parser: "ParseSANS"}
}
