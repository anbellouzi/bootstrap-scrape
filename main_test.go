package main

import (
	"testing"
)

func TestPostAPI(t *testing.T) {
	test_name := "test component name"
	test_html := "<h1>hello</h1>"
	expected := "200 OK"
	resp := post_to_API(test_name, test_html)
	if resp.Status != expected {
		t.Errorf("post_to_api() test returned an unexpected result: got %v want %v", resp.Status, expected)
	}

}

func TestDeleteAPI(t *testing.T) {
	expected := "200 OK"
	resp := removeAll_APIData()
	if resp.Status != expected {
		t.Errorf("deleteAll_from_api() test returned an unexpected result: got %v want %v", resp.Status, expected)
	}
}

func TestGetAPI(t *testing.T) {
	test_file_name := "test_file.txt"
	expected := "200 OK"
	resp := get_api_components(test_file_name)

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
