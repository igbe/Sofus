package config

import (
	"reflect"
	"testing"
)

type Org string
type Url string

func TestLoadConfig(t *testing.T) {
	// Configuration holds the urls of the various supported sites
	type Configuration struct {
		Orgs map[Org][]Url `mapstructure:"urlConfig"`
	}

	type configParams struct {
		path      string
		fileNAme  string
		extension string
		conf      *Configuration
	}

	cp := configParams{
		path:      "test_data/",
		fileNAme:  "config",
		extension: "yaml",
		conf:      &Configuration{},
	}

	want := Configuration{Orgs: map[Org][]Url{"sans": {"https://www.example.com"}}}

	err := LoadConfig(cp.path, cp.fileNAme, cp.extension, cp.conf)
	if err != nil {
		t.Fatalf("TestLoadConfig: %#v", err)
	}
	got := *cp.conf

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TestLoadConfig: Expected %#v but got %#v", want, got)
	}

}
