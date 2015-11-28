package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	words := strings.Fields(s)
	fmt.Printf("words : %v\n", words)
	for w := 0; w < len(words); w++{
		fmt.Printf("word : %v\n", words[w])
		fmt.Printf("len : %v\n", len(words[w]))
		(wordMap[words[w]])++
	}
	
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
