package lib

import (
	"fmt"
	"os"
)

// Function that get a file content
func GetFileContent(fileName string) (string, bool) {
	thinkertoy := false
	if fileName == "assets/thinkertoy.txt" {
		thinkertoy = true
	}
	fileName = "asciiart/" + fileName
	b, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	content := string(b)
	return content, thinkertoy
}
