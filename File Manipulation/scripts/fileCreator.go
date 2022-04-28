package scripts

import "strconv"

//GenerateFileName returns a string in form of a Name for the output.csv file numeric ordered
func GenerateFileName(chunkNumber int) (fileName string) {
	fileName = "output"
	fileName = fileName + strconv.Itoa(chunkNumber) + ".csv"
	return fileName
}
