package main

import (
	"fmt"
	"io"
	"net/http"

	config "github.com/igbe/Sofus/pkg"
)

type Org string
type Url string

// Configuration holds the urls of the various supported sites
type Configuration struct {
	Orgs map[Org][]Url `mapstructure:"urlConfig"`
}

// Configuration holds the configs needed to run this app and is loaded from
// the user specified config file.
type configParams struct {
	path      string
	fileNAme  string
	extension string
	conf      *Configuration
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

func processSansPages(ulrs []Url) {
	fmt.Println(ulrs)
	// resp, err := fetchPage(URL.sans.detection)
	// if err != nil {
	// 	fmt.Printf("Error retrieving page, %v", err)
	// }

	// dom, err := goquery.NewDocumentFromReader(strings.NewReader(resp))
	// if err != nil {
	// 	fmt.Printf("Error retrieving dom, %v", err)
	// }

	// dom.Find(".rr_paper").Each(func(index int, element *goquery.Selection) {
	// 	fmt.Println(index, element.Text())
	// 	//href, _ := element.Attr("href")

	// })
}

func main() {
	cp := configParams{
		path:      "../configs",
		fileNAme:  "config",
		extension: "yaml",
		conf:      &Configuration{},
	}
	err := config.LoadConfig(cp.path, cp.fileNAme, cp.extension, cp.conf)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	// Process each org/site by calling the appropriate extractor
	for org, url := range (*cp.conf).Orgs {
		switch org {
		case "sans":
			processSansPages(url)
		default:
			fmt.Println("No Extraction function have been written for this Org/site")
		}
	}

}
