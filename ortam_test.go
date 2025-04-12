package ortam

import (
	"net/url"
	"os"
	"testing"
)

type SubConfig struct {
	UrlConf *url.URL
}

type Config struct {
	StringConf string
	UintConf   uint16
	IntConf    int
	FloatConf  float64
	BoolConf   bool
	StructConf SubConfig `ortam:"SUB"`
}

func badValue(t *testing.T, name string, val any) {
	t.Errorf("%s configuration option does not have the expected value: %v", name, val)
}

func TestLoad(t *testing.T) {
	var config Config = Config{
		StringConf: "change me",
		UintConf:   42,
	}

	os.Setenv("TEST_STRING_CONF", "hello world")
	os.Setenv("TEST_INT_CONF", "-42")
	os.Setenv("TEST_FLOAT_CONF", "42.713")
	os.Setenv("TEST_BOOL_CONF", "true")
	os.Setenv("TEST_SUB_URL_CONF", "http://localhost:8080")

	if err := Load(&config, "TEST"); err != nil {
		t.Errorf("load failed: %s", err.Error())
	}

	if config.StringConf != "hello world" {
		badValue(t, "string", config.StringConf)
	}

	if config.IntConf != -42 {
		badValue(t, "int", config.IntConf)
	}

	if config.UintConf != 42 {
		badValue(t, "uint", config.UintConf)
	}

	if !config.BoolConf {
		badValue(t, "bool", config.BoolConf)
	}

	if config.FloatConf != 42.713 {
		badValue(t, "float", config.FloatConf)
	}

	if config.StructConf.UrlConf == nil {
		t.Errorf("url configuration option is nil")
	} else if config.StructConf.UrlConf.Host != "localhost:8080" {
		badValue(t, "url", config.StructConf.UrlConf.String())
	}
}
