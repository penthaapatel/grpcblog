package serializer

import (
	"encoding/json"
	"fmt"
	"grpcblog/blog"
	"io/ioutil"
	"os"
)

type BlogStruct struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

//checkFile checks if a file exists
func checkFile(filename string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("cannot open file: %w", err)
	}

	defer f.Close()
	return nil
}

// WriteProtobufToJSONFile writes protocol buffer message to JSON file
func WriteProtobufToJSONFile(b *blog.Blog, filename string) error {
	err := checkFile(filename)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	structdata := []BlogStruct{}

	json.Unmarshal(file, &structdata)

	newStruct := &BlogStruct{
		Title: b.Title,
		Body:  b.Body,
	}
	structdata = append(structdata, *newStruct)

	// Preparing the data to be marshalled and written.
	dataBytes, err := json.Marshal(structdata)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	err = ioutil.WriteFile(filename, dataBytes, 0644)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
