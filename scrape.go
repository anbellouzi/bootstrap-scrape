package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type component struct {
	Name string `json:"name"`
	Html string `json:"html"`
}

// write data to specific file
func writeFile(file []byte, filename string) {
	this := ioutil.WriteFile(filename, file, 0644)
	if err := this; err != nil {
		panic(err)
	}
	fmt.Println("Saved to", filename, "✅")
}

// serialize components and replace unwaned chars
func serializeJSON(foo []component, filename string) {
	fmt.Print("Serializing data... ")
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(foo)
	res := bytes.ReplaceAll(bf.Bytes(), []byte("\\n"), []byte(""))
	res = bytes.ReplaceAll(res, []byte("\\"), []byte(""))
	fmt.Println("✅")
	writeFile(res, filename)
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

// part of code adapted from respond by metalim on stackoverflow
// https://stackoverflow.com/questions/19253469/make-a-url-encoded-post-request-using-http-newrequest
func post_to_API(name, html string) {
	apiUrl := "https://bootstrap-api.herokuapp.com/components/add/component"

	data := url.Values{}
	data.Set("name", name)
	data.Set("html", html)

	client := &http.Client{}
	r, err := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		fmt.Println(err)
		return
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(r)
	fmt.Println(resp.Status, name, "Saved! ✅")
}

// removes all components saved inside API
func removeAll_APIData() {
	apiUrl := "https://bootstrap-api.herokuapp.com/components/delete_all"

	client := &http.Client{}
	r, err := http.NewRequest("DELETE", apiUrl, nil) // URL-encoded payload
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, _ := client.Do(r)
	fmt.Println(resp.Status, "Removed all components from API! ✅")
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func scrapeData(toAPI, toFile bool, dataFile, filename string) {
	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML(".span9", func(e *colly.HTMLElement) {

		var components []component
		var element component
		data, _ := readLines(dataFile)
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
				if toAPI == true {
					post_to_API(element.Name, element.Html)
				}
				counter = 0
				components = append(components, element)
			}
		}
		if toFile == true {
			serializeJSON(components, filename)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://getbootstrap.com/2.3.2/components.html")
}

func main() {
	filename := flag.String("filename", "output.json", "name of .json or .txt file you want to save components to")
	toApi := flag.Bool("toApi", false, "true or false do you want to save components to API?")
	toFile := flag.Bool("toFile", true, "true or false do you want to save components to a file?")
	dataFile := flag.String("data", "components.txt", ".txt file that contains html selectors you want to scrape")
	removeAll := flag.Bool("remove", false, "true or false do you want to remove all data from API?")
	flag.Parse()

	if *removeAll == true {
		removeAll_APIData()
	} else {
		scrapeData(*toApi, *toFile, *dataFile, *filename)
	}

}
