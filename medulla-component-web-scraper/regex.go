package main

import (
	"regexp"
	"strings"
)

const (
	RegexUtilHrefMatch = "href=\"(.+?)\""
)

func GetMatches(content string, regex string) [][]byte {
	re := regexp.MustCompile(regex)
	matches := re.FindAll([]byte(content), -1)
	return matches
}

func GetLinks(content string) [][]byte {
	var newMatches [][]byte
	matches := GetMatches(content, RegexUtilHrefMatch)
	for i := 0; i < len(matches); i++ {
		cleanStr := strings.ReplaceAll(string(matches[i]), "href=\"", "")
		cleanStr = strings.ReplaceAll(cleanStr, "\"", "")
		newMatches = append(newMatches, []byte(cleanStr))
	}
	return newMatches
}
