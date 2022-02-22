package main

import (
	"fmt"
	"log"
	"testing"
)

func TestGetMatches(t *testing.T) {
	html := GetHtmlAt(ExternalComponentsLibrariesPrimer)
	matches := GetMatches(html, RegexUtilHrefMatch)
	if len(matches) == 0 {
		t.Fatal("Failed to get matches")
	}
	fmt.Println("\r\n\r\n\r\n\r\nTesting ability to get href from html")
	for i := 0; i < len(matches); i++ {
		fmt.Printf(""+
			"\r\n"+
			"match: %s"+
			"\r\n", matches[i])
	}
	fmt.Println("Test Complete")
}

func TestGetLinks(t *testing.T) {
	log.Println("TestGetLinks")
	html := GetHtmlAt(ExternalComponentsLibrariesPrimer)
	matches := GetLinks(html)
	if len(matches) == 0 {
		t.Fatal("Failed to get matches")
	}
	fmt.Println("\r\n\r\n\r\n\r\nTesting ability to get href from html")
	for i := 0; i < len(matches); i++ {
		fmt.Printf(""+
			"\r\n"+
			"match: %s"+
			"\r\n", matches[i])
	}
	fmt.Println("Test Complete")
}
