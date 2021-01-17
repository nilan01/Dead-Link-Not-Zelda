package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	"mvdan.cc/xurls"
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

// Extract links from entire html tags
func extractLinks(stringArray []string) {
	rxRelaxed := xurls.Relaxed()
	for i := 0; i < len(stringArray); i++ {
		stringArray[i] = rxRelaxed.FindString(stringArray[i])
	}
}

// Send GET request to all links in array
func httpRequestLinks(stringArray []string) {
	//colorGreen := "\033[32m"
	c1 := color.New(color.FgHiGreen).Add()
	c2 := color.New(color.FgHiRed).Add()
	c3 := color.New(color.FgHiBlack).Add()
	for i := 0; i < len(stringArray); i++ {
		resp, err := http.Head(stringArray[i])
		if err != nil {
			c3.Printf("%-10v", "UNKNOWN")
			c3.Println(stringArray[i])
		} else {
			if resp.StatusCode == 200 {
				c1.Printf("%-10v", "GOOD")
				c1.Println(stringArray[i])
			} else if resp.StatusCode == 400 || resp.StatusCode == 404 {
				c2.Printf("%-10v", "BAD")
				c2.Println(stringArray[i])
			}
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

	// Fill the made array of strings with content which is an array of bytes.
	fillStringArray(stringArray, content)

	// Extract all links from the array.
	extractLinks(stringArray)

	// Send GET Requests to array of links
	httpRequestLinks(stringArray)

}
