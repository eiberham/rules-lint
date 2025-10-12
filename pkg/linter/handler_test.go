package linter

import (
	"os"
	"testing"
)

type TestLineHandler struct {
	BaseHandler
}

func (tlh *TestLineHandler) Handle(rule Rule, cfg *Config, file *os.File) []string {
	return []string{"Runs line handler"}
}

type TestFileHandler struct {
	BaseHandler
}

func (tfh *TestFileHandler) Handle(rule Rule, cfg *Config, file *os.File) []string {
	return []string{"Runs file handler"}
}

func TestHandler(t *testing.T) {

	rule := &LineRule{}

	cfg := &Config{}
	expected := []string{"Runs line handler"}

	processor := &BaseHandler{}
	processor.
		Next(&TestLineHandler{}).
		Next(&TestFileHandler{})

	received := processor.Handle(rule, cfg, nil)

	if !isEqual(received, expected) {
		t.Errorf("Expected %v, but got %v", expected, received)
	}

}
