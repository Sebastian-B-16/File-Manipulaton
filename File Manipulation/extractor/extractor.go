package extractor

import (
	"bufio"
	"log"
	"os"
)

//ReadFile function return the content of the input.txt file
func ReadFile() (scanner *bufio.Reader, err error) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Print("Error trying to open the txt file: ", err)
		return scanner, err
	}
	scanner = bufio.NewReader(f)
	return scanner, nil
}

//ReadHeader function return the header of the input.txt file
func ReadHeader(scanner *bufio.Reader) (header string, err error) {
	header, err = scanner.ReadString('\n')
	if err != nil {
		log.Print("Error Trying to read the header: ", err)
		return "", err
	}
	return header, nil
}
