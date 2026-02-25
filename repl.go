package main

import (
	"strings"
)

func cleanInput(text string) []string {
	newString := strings.ToLower(text)
	res := strings.Fields(newString)
	return res
}