package db

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"school-system/cmd/repository/db/file_handler"
	dbutils "school-system/cmd/repository/db/utils"
	"strings"

	"strconv"
)

func createDBFile(filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE, 0644)

	if err != nil {
		slog.Error(fmt.Sprintf("creating file %v\n", filename))
		return err
	}

	file.Close()
	return nil
}

func initDB() error {
	file, errorReadingFile := os.Open(getDBFilepath())

	if errorReadingFile != nil {

		errorCreatingFile := createDBFile(getDBFilepath())

		if errorCreatingFile != nil {
			slog.Error(fmt.Sprintf("error %v initializing DB", errorCreatingFile.Error()))
			return errorCreatingFile
		}
	}
	file.Close()
	return nil
}

func getDBFileLocation() string {
	dbfile := "/db.txt"

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		slog.Warn("Could not determine the DB file path location. Returning default.")
		return "cmd/server/db" + dbfile
	}
	workDir := filepath.Dir(file)
	dbFilename := workDir + dbfile

	slog.Info("DB file located", "DBPath", dbFilename)
	return dbFilename
}

var dbFilename string

func getDBFilepath() string {
	if dbFilename == "" {
		dbFilename = getDBFileLocation()
	}

	return dbFilename
}

type DB struct{}

type Document interface {
	ToString() string
}

func GetDB() *DB {
	if err := initDB(); err == nil {
		return &DB{}
	}

	return nil
}

var Instance = GetDB()

// Fake DB: All is based on files
func (d *DB) Insert(data interface{}) error {
	dbFile, err := file_handler.OpenFileWithPerm(getDBFilepath(), os.O_APPEND|os.O_WRONLY)
	if err != nil {
		return err
	}
	defer dbFile.Close()
	dataString := dbutils.ConvertStructToString(data)
	file_handler.AppendToFile(dbFile, dataString)
	return nil
}

func (d *DB) InsertAll(dataList []Document) error {
	dbFile, err := file_handler.OpenFileWithPerm(getDBFilepath(), os.O_APPEND|os.O_WRONLY)
	if err != nil {
		return err
	}
	defer dbFile.Close()
	var builder strings.Builder
	for _, v := range dataList {
		dataString := dbutils.ConvertStructToString(v)
		builder.WriteString(fmt.Sprintf("%s\n", dataString))
	}
	file_handler.AppendToFile(dbFile, builder.String())
	return nil
}

func (d *DB) Update(id int, data interface{}) error {
	dbFile, err := file_handler.OpenFileWithPerm(getDBFilepath(), os.O_RDWR)
	if err != nil {
		return err
	}
	defer dbFile.Close()
	dataString := dbutils.ConvertStructToString(data)
	err = file_handler.UpdateFileEntry(dbFile, strconv.Itoa(id), dataString)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) GetAll() ([]string, error) {
	dbFile, err := os.OpenFile(getDBFilepath(), os.O_RDWR, 0644)
	if err != nil {
		slog.Error(err.Error())
		return []string{}, err
	}
	dbFileContent := file_handler.GetFileContent(dbFile)
	// Remove any empty line that may exists.
	file_handler.OverrideFileContent(dbFile, dbFileContent)
	return dbFileContent, nil
}

func (d *DB) GetByID(id int) (string, error) {
	dbFile, err := file_handler.OpenFileWithPerm(getDBFilepath(), os.O_RDONLY)
	if err != nil {
		return "", fmt.Errorf("error %v trying get content of ID %v", err, id)
	}
	defer dbFile.Close()
	content, err := file_handler.GetFileEntryByPrefix(dbFile, strconv.Itoa(id))
	if err != nil {
		return "", err
	}
	return content, nil
}

func (d *DB) Delete(id int) error {
	dbFile, err := file_handler.OpenFileWithPerm(getDBFilepath(), os.O_RDWR)
	if err != nil {
		return err
	}
	defer dbFile.Close()
	file_handler.RemoveFileEntry(dbFile, strconv.Itoa(id))
	return nil
}

func (d *DB) Clear() error {
	dbFile, err := file_handler.OpenFileWithPerm(getDBFilepath(), os.O_TRUNC)
	if err != nil {
		return err
	}
	defer dbFile.Close()
	file_handler.ClearFileContent(dbFile)
	return nil
}
