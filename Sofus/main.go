package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type SANS struct {
	detection string
}

type url struct {
	sans SANS
}

var (
	URL = url{
		sans: SANS{
			detection: "https://www.sans.org/reading-room/whitepapers/detection/",
		},
	}
)

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

func main() {
	resp, err := fetchPage(URL.sans.detection)
	if err != nil {
		fmt.Printf("Error retrieving page, %v", err)
	}

	//fmt.Println(resp)

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(resp))
	if err != nil {
		//log.Fatalln(err)
		fmt.Printf("Error retrieving dom, %v", err)
	}

	dom.Find(".rr_paper").Each(func(index int, element *goquery.Selection) {
		//fmt.Println(i, element.Text())
		//href, _ := element.Attr("href")

	})
}

// // convert the return response to string
// b, err := io.ReadAll(resp.Body)
// if err != nil {
// 	return "", fmt.Errorf("unable to read the http response body into io: %v", resp.Body)
// }

// return string(b), nil
