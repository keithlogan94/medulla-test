package main

import (
	"strings"
	"testing"
)

func TestGetHtmlAt(t *testing.T) {
	var html string = GetHtmlAt("https://www.google.com")
	var checkContains string = "<!doctype html>"
	if !strings.Contains(html, checkContains) {
		t.Fatalf("html does not contain %s", checkContains)
	}
}
