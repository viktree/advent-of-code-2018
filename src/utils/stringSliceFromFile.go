package utils

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func StringSliceFromFile(relativeFilePath string) []string {
	absPath, _ := filepath.Abs(relativeFilePath)

	rawFileContent, err := ioutil.ReadFile(absPath)
	if err != nil {
		fmt.Print(err)
	}

	fileContent := string(rawFileContent)
	linesOfFile := strings.Split(fileContent, "\n")

	return linesOfFile
}
