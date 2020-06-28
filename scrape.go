package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

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

func serializeJSON(foo []component) {
	fmt.Println("Serializing Data")
	fooJSON, _ := json.Marshal(foo)

	writeFile(fooJSON)
	fmt.Println("Serializing Complete ")
	fmt.Println(string(fooJSON))
}

func clear(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
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

	c.OnHTML(".span9", func(e *colly.HTMLElement) {

		var components []component
		var element component
		data, _ := readLines("components.txt")

		counter := 0
		for _, line := range data {
			if line != "" {
				if counter%2 == 0 {
					element.Name = e.ChildText(line)
				} else {
					element.Html = e.ChildText(line)
				}
				counter = counter + 1
			} else {
				counter = 0
				components = append(components, element)
			}
		}

		serializeJSON(components)

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://getbootstrap.com/2.3.2/components.html")

	// serialize data to json and write it to file

	// fmt.Println(element.Name, element.Html)

}
