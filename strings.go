package carbot

import "regexp"

// removeMultipleSpaces removes multiple spaces to single space
func removeMultipleSpaces(s string) string {
	// use regexp to remove multiple spaces to single space
	re := regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, " ")
	return s
}

// removeSpacesBetweenDigits removes spaces between digits
func removeSpacesBetweenDigits(s string) string {
	// use regexp to remove spaces between digits
	re := regexp.MustCompile(`(\d)\s+(\d)`)
	s = re.ReplaceAllString(s, "$1$2")
	return s
}
