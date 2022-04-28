package main

import (
	"File_Manipulation/extractor"
	"File_Manipulation/loader"
	"fmt"
	"log"
)

func main() {

	var chunkSize int
	//Initialize the chunk size
	fmt.Print("Choose the Chunk Size: ")
	_, err := fmt.Scanf("%d", &chunkSize)
	if err != nil {
		log.Print("Error trying to assign a chunkSize parameter")
	}
	//Extract the data
	scanner, err := extractor.ReadFile()
	if err != nil {
		log.Print("Error trying to scan the input file")
	}
	//Extract the header
	header, err := extractor.ReadHeader(scanner)
	if err != nil {
		log.Print("Error trying to extract the header")
	}
	//Load the data into new output files
	loader.LoadFiles(scanner, header, chunkSize)
}
