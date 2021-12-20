package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	config "github.com/igbe/Sofus/pkg"
	log "github.com/sirupsen/logrus"
)

// For logging
var contextLogger *log.Entry

// Configuration holds the configs needed to run this app and is loaded from
// the user specified config file.
type configParams struct {
	path      string
	fileName  string
	extension string
}

type Sites struct {
	url        string
	ParserName string
}

// Initialize important things that need to run before main
func init() {
	// Set formatter to have th elog be formted in JSON
	log.SetFormatter(&log.JSONFormatter{})

	// Get the current home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// Write all  log output to a log file in the home dir
	// If the file doesn't exist, create it or append to the file
	logFile, err := os.OpenFile(filepath.Join(homeDir, ".sofus.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	// Initialize constant fields that would be seen on all log lines.
	contextLogger = log.WithFields(log.Fields{
		"Name": "Sofus",
	})
}

// fetchSites parses the items in the conf config and returns the url and the name of teh function to parse the url contents.
func fetchSites(conf interface{}) []Sites {
	var sites []Sites
	// Note: Type assertion are performed on the conf interface before accessing the Orgs key.
	for _, urlWithFunc := range conf.(config.Configuration).Orgs {
		res := strings.Split(strings.ReplaceAll(urlWithFunc, " ", ""), ",")
		sites = append(sites, Sites{url: res[0], ParserName: res[1]})
	}
	return sites
}

func main() {
	cp := configParams{
		path:      "configs",
		fileName:  "config",
		extension: "yaml",
	}

	contextLogger.Info("Loading configs from config.yaml file")
	conf, err := config.LoadConfig(cp.path, cp.fileName, cp.extension)
	if err != nil {
		contextLogger.Error(err)
	}

	// Channel to hold teh return values from the parsers.
	cOut := make(chan PageOutput)

	sites := fetchSites(conf)

	// Process each org/site by calling the appropriate extractor
	for _, s := range sites {
		go func(url string, parserFunc string) {
			cOut <- Parsers[parserFunc](url)
		}(s.url, s.ParserName)
	}

	// Collect results:
	results := make([]PageOutput, len(sites))

	for i := range results {
		out := <-cOut

		// Check if the parser returned empty result
		if len(out.articles) == 0 {
			contextLogger.Warnf(fmt.Sprintf("%s: returned empty article list for configured url", out.parser))
		}

		results[i] = out
	}
	// For now Temporarily print output to console
	fmt.Printf("Results: %+v\n", results)

}
