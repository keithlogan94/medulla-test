package main

import (
	"log"
	"strings"
)

func PrintList(items [][]byte) {
	for i := 0; i < len(items); i++ {
		log.Print("\r\n\r\n\r\n\r\n\r\n\r\n")
		log.Printf("item: %s", items[i])
		log.Print("\r\n\r\n\r\n\r\n\r\n\r\n")
	}
}

func FilterList(items [][]byte, containsString string) [][]byte {
	var list [][]byte
	for i := 0; i < len(items); i++ {
		if strings.Contains(string(items[i]), containsString) {
			list = append(list, items[i])
		}
	}
	return list
}

func PrependListItemsWith(items [][]byte, prependWith []byte) [][]byte {
	var list [][]byte
	for i := 0; i < len(items); i++ {
		list = append(list, []byte(string(prependWith)+string(items[i])))
	}
	return list
}
