package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

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

// Kicks off the tasks assigned to the various workers which includes
// visiting the sites and retrieving the pages and then parsing the
// title and other useful values.
func worker(url string, parserFunc string) {
	fmt.Printf("Processing url:%s with function: %s\n", url, parserFunc)
	time.Sleep(3 * time.Second)
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

	// Write all  log output to a log file in teh home dir
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

func main() {
	var wg sync.WaitGroup

	cp := configParams{
		path:      "../configs",
		fileName:  "config",
		extension: "yaml",
	}

	contextLogger.Info("Loading configs from config.yaml file")
	conf, err := config.LoadConfig(cp.path, cp.fileName, cp.extension)
	if err != nil {
		contextLogger.Error(err)
	}

	//Process each org/site by calling the appropriate extractor
	//Note: we are performing type assertion on the conf interface before
	//accessing the Orgs key.
	for _, urlWithFunc := range conf.(config.Configuration).Orgs {
		res := strings.Split(strings.ReplaceAll(urlWithFunc, " ", ""), ",")
		url := res[0]
		parserFunc := res[1]

		wg.Add(1)
		go func(url string, parserFunc string) {
			defer wg.Done()
			worker(url, parserFunc)
		}(url, parserFunc)
	}
	wg.Wait()

}
