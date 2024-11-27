package file_handler

import (
	"bufio"
	"fmt"
	"log"
	"os"
	utils "school-system/cmd/app/utils"

	"strings"
)

func OpenFileWithPerm(filename string, flag int) (*os.File, error) {
	file, err := os.OpenFile(filename, flag, 0644)

	if err != nil {
		log.Fatalf("ERROR %v opening file %v", err, filename)
		return nil, err
	}

	return file, nil
}

func PrintFileContent(file *os.File) {
	fileContent := GetFileContent(file)

	for _, v := range fileContent {
		fmt.Println(v)
	}
}

func GetFileContent(file *os.File) []string {
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	var content []string

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			content = append(content, line)
		}
	}
	return content
}

func GetFileEntryByPrefix(file *os.File, prefix string) string {
	fileContent := GetFileContent(file)

	for _, v := range fileContent {
		vPrefix := strings.Split(v, " ")[0]
		if vPrefix == prefix {
			return v
		}
	}

	return ""
}

func UpdateFileEntry(file *os.File, entryPrefix, updatedEntry string) {
	fileContent := GetFileContent(file)
	var newContent []string

	for _, v := range fileContent {

		vPrefix := strings.Split(v, " ")[0]
		if vPrefix == entryPrefix {
			newContent = append(newContent, updatedEntry)
		} else if v == "" {
			continue
		} else {
			newContent = append(newContent, v)
		}
	}

	OverrideFileContent(file, newContent)
}

func RemoveFileEntry(file *os.File, entryPrefix any) {
	fileContent := GetFileContent(file)
	var newContent []string

	for _, v := range fileContent {
		vPrefix := strings.Split(v, " ")[0]
		if vPrefix == entryPrefix || v == "" {
			continue
		} else {
			newContent = append(newContent, v)
		}
	}

	OverrideFileContent(file, newContent)
}

func OverrideFileContent(file *os.File, content []string) {
	file.Truncate(0)

	file.Seek(0, 0)

	utils.SortSliceStringByID(content, " ")

	for _, v := range content {
		file.WriteString(fmt.Sprintf("%s\n", v))
	}
}

func AppendToFile(file *os.File, content string) {
	file.WriteString(fmt.Sprintf("\n%s", content))
}

func ClearFileContent(file *os.File) {
	file.Truncate(0)
}
