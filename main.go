package main

import (
	"fmt"
	"log"
	"os"

	"rsc.io/quote"
)

const (
	outputFileName  = "output.txt"
	outputFileFlags = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	outputFilePerm  = 0644
)

func get_text_file() string {
	// This function is a placeholder for the actual implementation
	// that would read from a text file.

	content, err := os.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(content)

}

func main() {
	fmt.Println("Welcome to the Simple Go ETL!")

	extractedText := get_text_file()
	fmt.Println("Data extracted from text.txt")

	quote := quote.Go()
	transformedData := fmt.Sprintf("Extracted: %s\nQuote: %s", extractedText, quote)

	fmt.Printf("Adding this data: %s\n", transformedData)

	// add extracted text and quote as a new line to a new text file
	file, err := os.OpenFile(outputFileName, outputFileFlags, outputFilePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString(transformedData + "\n"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data transformed and loaded into output.txt")

}
