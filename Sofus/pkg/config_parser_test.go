package config

import (
	"reflect"
	"testing"
)

type Org string
type Url string

func TestLoadConfig(t *testing.T) {

	type configParams struct {
		path      string
		fileName  string
		extension string
	}

	cp := configParams{
		path:      "test_data/",
		fileName:  "config",
		extension: "yaml",
	}

	want := Configuration{
		Orgs:    []string{"https://www.example.com"},
		Workers: map[string]string{"retryinterval": "2"},
	}

	got, err := LoadConfig(cp.path, cp.fileName, cp.extension)
	if err != nil {
		t.Fatalf("TestLoadConfig: %#v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TestLoadConfig: Expected %#v but got %#v", want, got)
	}

}
