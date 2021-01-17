package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Count number of elements we will need in the string array.
func determineArraySize(content []byte) int {
	count := 0
	for i := 0; i < len(content); i++ {
		if string(content[i]) == "\n" {
			count++
		}
	}
	return count
}

// For each element in the stringArray, add a character (stored in context) until a character is equal to a new-line.
// Do this n amount of times for every element in stringArray
func fillStringArray(emptyArray []string, content []byte) {
	characterCount := 0
	for k := 0; k < len(emptyArray); k++ {
		for x := characterCount; string(content[x]) != "\n"; x++ {
			emptyArray[k] += string(content[x])
			characterCount = x + 2
		}
	}
}

func main() {
	// Read file in local directory - store contents in content.
	fileName := os.Args[1]
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// Make array of strings with specific number of elements.
	stringArray := make([]string, determineArraySize(content))

	// Fill the made array of strings with content which is an array of bytes
	fillStringArray(stringArray, content)

	fmt.Println(stringArray[2])
}