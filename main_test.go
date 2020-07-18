package main

import (
	"testing"
)

func TestPostAPI(t *testing.T) {
	test_name := "test component name"
	test_html := "<h1>hello</h1>"
	expected := "200 OK"
	resp := post_to_API(test_name, test_html, false)
	if resp.Status != expected {
		t.Errorf("post_to_api() test returned an unexpected result: got %v want %v", resp.Status, expected)
	}

}

func TestDeleteAPI(t *testing.T) {
	expected := "200 OK"
	resp := removeAll_APIData(false)
	if resp.Status != expected {
		t.Errorf("deleteAll_from_api() test returned an unexpected result: got %v want %v", resp.Status, expected)
	}
}

func TestGetAPI(t *testing.T) {
	test_file_name := "test_file.txt"
	expected := "200 OK"
	resp := get_api_components(test_file_name, false)
	if resp.Status != expected {
		t.Errorf("post_to_api() test returned an unexpected result: got %v want %v", resp.Status, expected)
	}
}

func TestWriteFile(t *testing.T) {
	test_file_name := "test_file.txt"
	if fileExists(test_file_name) == false {
		t.Errorf("TestGetAPI() test did not create proper file: want %v", test_file_name)
	}
}

func TestReadLines(t *testing.T) {
	test_file_name := "components.txt"
	_, err := readLines(test_file_name)
	if err != nil {
		t.Errorf("TestReadLines() test did not find the data file: err %v", err)
	}
}

func BenchmarkPostAPI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		post_to_API("test component name", "<h1>hello</h1>", false)
	}
}

func BenchmarkReadAPI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		get_api_components("readAPI-test.txt", false)
	}
}

func BenchmarkScrapeSingleToNON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dataFile := "single_component.txt"
		scrapeData(false, false, dataFile, "scrape-to-non.txt", false)
	}
}

func BenchmarkScrapeSingleToFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dataFile := "single_component.txt"
		scrapeData(false, true, dataFile, "scrape-tofile-test.txt", false)
	}
}

func BenchmarkScrapeSingleToAPI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dataFile := "single_component.txt"
		scrapeData(true, false, dataFile, "singleToAPI.txt", false)
	}
}

func BenchmarkScrapeSingleToALL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dataFile := "single_component.txt"
		scrapeData(true, true, dataFile, "single-toAll-test.txt", false)
	}
}

func BenchmarkScrapeAllToNON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dataFile := "components.txt"
		scrapeData(false, false, dataFile, "ScrapeAllToNON.txt", false)
	}
}

func BenchmarkScrapeAllToFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dataFile := "components.txt"
		scrapeData(false, true, dataFile, "AllToFile.txt", false)
	}
}

func BenchmarkDeleteAllAPI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeAll_APIData(false)
	}
}

func BenchmarkScrapeAllToAPI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dataFile := "components.txt"
		scrapeData(true, false, dataFile, "AllToAPI.txt", false)
	}
}
