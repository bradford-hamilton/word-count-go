package main

import (
	wc "github.com/bradford-hamilton/practice-1/test"
	"strings"
)

func WordCount(s string) map[string]int {
	myArr := strings.Fields(s)
	myMap := make(map[string]int)
	for i := 0; i < len(myArr); i++ {
		myMap[myArr[i]] += 1
	}

	return myMap
}

func main() {
	wc.Test(WordCount)
}
