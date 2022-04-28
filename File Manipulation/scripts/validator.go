package scripts

import "strings"

//ValidateLine function will check if the line data is fully completed and return a boolean by case.
func ValidateLine(line string) bool {
	sep := strings.Split(line, "|")
	for i := 0; i < len(sep); i++ {
		if sep[i] == "" {
			return false
		}
	}
	return true
}
