package db

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"reflect"
	"school-system/cmd/app/Utils/file_handler"
	"strconv"
	"strings"
)

var dbFilename = "cmd/app/db/students.txt"

// Fake DB: All is based on files
func Init() {
	file, errorReadingFile := os.Open(dbFilename)

	if errorReadingFile != nil {
		errorCreatingFile := createDBFile(dbFilename)

		if errorCreatingFile != nil {
			log.Fatal(errorCreatingFile)
		}
	}

	file.Close()
}

func Insert(data interface{}) bool {
	dbFile := file_handler.OpenFileWithPerm(dbFilename, os.O_APPEND|os.O_WRONLY)

	if dbFile != nil {
		defer dbFile.Close()
		dataString := convertStructToString(data)
		file_handler.AppendToFile(dbFile, dataString)
		return true
	} else {
		return false
	}
}

func Update(id int, data interface{}) bool {
	dbFile := file_handler.OpenFileWithPerm(dbFilename, os.O_RDWR)

	if dbFile != nil {
		defer dbFile.Close()
		dataString := convertStructToString(data)
		fmt.Println(data)
		fmt.Println(dataString)
		file_handler.UpdateFileEntry(dbFile, strconv.Itoa(id), dataString)

		return true
	} else {
		return false
	}
}

func GetAll() []string {
	dbFile, err := os.OpenFile(dbFilename, os.O_RDWR, 0644)

	if err != nil {
		log.Fatal(err)
	}

	dbFileContent := file_handler.GetFileContent(dbFile)

	// Remove any empty line that may exists.
	file_handler.OverrideFileContent(dbFile, dbFileContent)

	return dbFileContent
}

func GetByID(id int) (string, error) {
	dbFile := file_handler.OpenFileWithPerm(dbFilename, os.O_RDONLY)

	if dbFile != nil {
		defer dbFile.Close()
		content := file_handler.GetFileEntryByPrefix(dbFile, strconv.Itoa(id))
		return content, nil
	} else {
		return "", fmt.Errorf("error trying get content of ID %v", id)
	}
}

func Delete(id int) bool {
	dbFile := file_handler.OpenFileWithPerm(dbFilename, os.O_RDWR)

	if dbFile != nil {
		defer dbFile.Close()
		file_handler.RemoveFileEntry(dbFile, strconv.Itoa(id))

		return true
	} else {
		return false
	}
}

func Clear() bool {
	dbFile := file_handler.OpenFileWithPerm(dbFilename, os.O_TRUNC)

	if dbFile != nil {
		defer dbFile.Close()

		file_handler.ClearFileContent(dbFile)
		return true
	} else {
		return false
	}
}

func convertStructToString(s interface{}) string {
	structValue := reflect.ValueOf(s)
	structType := reflect.TypeOf(s)

	var builder strings.Builder

	for i := 0; i < structType.NumField(); i++ {
		value := structValue.Field(i)
		if value.CanInt() || value.CanConvert(reflect.TypeOf(string(""))) {
			builder.WriteString(fmt.Sprintf("%v ", value))
		}

		if value.CanConvert(reflect.TypeOf([]int{})) {
			println("PODE")
			convertedValue := value.Convert(reflect.TypeOf([]int{}))
			result, ok := convertedValue.Interface().([]int)

			if ok {
				for _, v := range result {
					builder.WriteString(fmt.Sprintf("%v ", strconv.Itoa(v)))
				}
			}
		}
	}

	return builder.String()
}

func createDBFile(filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE, 0644)

	if err != nil {
		slog.Error(fmt.Sprintf("creating file %v\n", filename))
		return err
	}

	file.Close()

	return nil
}
