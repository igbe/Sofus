package main

import "fmt"

// func fetchPage(url string) (string, error) {
// 	// Make HTTP GET request
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != 200 {
// 		return "", fmt.Errorf("wrong status code: %d", resp.StatusCode)
// 	}

// 	// Convert the output to string
// 	b, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", fmt.Errorf("unable to read the http response body into io: %v", resp)
// 	}

// 	return string(b), nil
// }

// This paerser handles SANS pages
func ParseSANS(url string) {
	fmt.Println(url)
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
