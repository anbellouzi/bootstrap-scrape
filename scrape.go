package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gocolly/colly"
)

type component struct {
	Name string `json:"name"`
	Html string `json:"html"`
}

func writeFile(file []byte) {
	this := ioutil.WriteFile("output.json", file, 0644)
	if err := this; err != nil {
		panic(err)
	}
}

func serializeJSON(foo component) {
	fmt.Println("Serializing Data", foo)
	fooJSON, _ := json.Marshal(foo)
	writeFile(fooJSON)
	fmt.Print("Serializing Complete ")
	fmt.Println(string(fooJSON))
}

// returns data as a string from file
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	var element component

	c.OnHTML("#dropdowns > h2:nth-child(6)", func(e *colly.HTMLElement) {
		element.Name = e.Text
		// fmt.Print(e.Text)
	})

	// coin name
	c.OnHTML("#dropdowns > pre:nth-child(8)", func(e *colly.HTMLElement) {
		element.Html = e.Text
	})

	// coin name
	c.OnHTML(".span9", func(e *colly.HTMLElement) {
		// element.Html = e.Text
		// fmt.Println(e.ChildText("#buttonDropdowns > div.page-header > h1"))
		// fmt.Println(e.ChildText("#dropdowns > pre:nth-child(21)"))

		data, _ := readLines("components.txt")

		counter := 0
		for _, line := range data {

			if line != "" {
				if counter%2 == 0 {
					fmt.Println("Name:", e.ChildText(line))
				} else {
					fmt.Println("HTML: ", e.ChildText(line))
				}

				counter = counter + 1
			} else {
				counter = 0
				fmt.Println("__________________________________________")
			}

		}

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://getbootstrap.com/2.3.2/components.html")

	// serialize data to json and write it to file
	// serializeJSON(element)

	// fmt.Println(element.Name, element.Html)

}
