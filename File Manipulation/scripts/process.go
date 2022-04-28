package scripts

import (
	"strconv"
	"strings"
)

//ProcessLine Function will process the lines that is to be loaded in the output.csv files
//And placing the new designated Id for each line.
func ProcessLine(id int, line string) (newLine string) {
	formatLine := strings.Split(line, "|")
	formatLine[0] = strconv.Itoa(id)
	newLine = strings.Join(formatLine, "|")
	return newLine
}
