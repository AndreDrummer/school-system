package db

import (
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"school-system/cmd/app/utils/file_handler"
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
			slog.Error(fmt.Sprintf("error %v initializing DB", errorCreatingFile.Error()))
			return
		}
	}

	file.Close()
}

func Insert(data interface{}) (bool, error) {
	dbFile, err := file_handler.OpenFileWithPerm(dbFilename, os.O_APPEND|os.O_WRONLY)

	if err != nil {
		return false, err
	}

	defer dbFile.Close()
	dataString := convertStructToString(data)
	file_handler.AppendToFile(dbFile, dataString)

	return true, nil

}

func Update(id int, data interface{}) (bool, error) {
	dbFile, err := file_handler.OpenFileWithPerm(dbFilename, os.O_RDWR)

	if err != nil {
		return false, err
	}

	defer dbFile.Close()
	dataString := convertStructToString(data)
	fmt.Println(data)
	fmt.Println(dataString)
	file_handler.UpdateFileEntry(dbFile, strconv.Itoa(id), dataString)

	return true, nil

}

func GetAll() ([]string, error) {
	dbFile, err := os.OpenFile(dbFilename, os.O_RDWR, 0644)

	if err != nil {
		slog.Error(err.Error())
		return []string{}, err
	}

	dbFileContent := file_handler.GetFileContent(dbFile)

	// Remove any empty line that may exists.
	file_handler.OverrideFileContent(dbFile, dbFileContent)

	return dbFileContent, nil
}

func GetByID(id int) (string, error) {
	dbFile, err := file_handler.OpenFileWithPerm(dbFilename, os.O_RDONLY)

	if err != nil {
		return "", fmt.Errorf("error %v trying get content of ID %v", err, id)
	}

	defer dbFile.Close()
	content := file_handler.GetFileEntryByPrefix(dbFile, strconv.Itoa(id))
	return content, nil
}

func Delete(id int) (bool, error) {
	dbFile, err := file_handler.OpenFileWithPerm(dbFilename, os.O_RDWR)

	if err != nil {
		return false, err
	}

	defer dbFile.Close()
	file_handler.RemoveFileEntry(dbFile, strconv.Itoa(id))

	return true, nil
}

func Clear() (bool, error) {
	dbFile, err := file_handler.OpenFileWithPerm(dbFilename, os.O_TRUNC)

	if err != nil {
		return false, err
	}

	defer dbFile.Close()

	file_handler.ClearFileContent(dbFile)
	return true, nil
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
