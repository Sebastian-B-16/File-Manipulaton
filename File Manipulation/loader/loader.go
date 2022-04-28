package loader

import (
	"File_Manipulation/scripts"
	"bufio"
	"io"
	"os"
	"strings"
)

//LoadFiles function will Load the data into newly created output.csv Files
func LoadFiles(scanner *bufio.Reader, header string, chunkSize int) {
	chunkNumber := 1
	valid := true
	for {
		f, err := os.Create(scripts.GenerateFileName(chunkNumber))
		if err != nil {
			return
		}
		//Applying the header to all the output files
		io.Copy(f, strings.NewReader(header))
		id := 1
		if !LoadData(f, scanner, chunkSize, id, valid) {
			break
		}
		chunkNumber++
	}
}

//LoadData Function will load the Data into its designed output.csv File
//By making calls to ValidateLine and ProcessLine scripts functions
//And return a boolean eoF variable in case we reached the end of File.
func LoadData(f *os.File, scanner *bufio.Reader, chunkSize int, id int, valid bool) (eoF bool) {
	for i := 0; i < chunkSize; i++ {
		line, err := scanner.ReadString('\n')
		if err != nil { //EOF Handler
			valid = false
			break
		}
		if scripts.ValidateLine(line) {
			newLine := scripts.ProcessLine(id, line)
			id++
			io.Copy(f, strings.NewReader(newLine))

		} else {
			i = i - 1
		}
	}
	eoF = valid
	return eoF
}
