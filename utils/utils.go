package utils

import (
	"fmt"
	"go-initializer/consts"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

func RandomString(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, n)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)

}

func HasElem(s interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(s)

	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {

			// XXX - panics if slice element points to an unexported struct field
			// see https://golang.org/pkg/reflect/#Value.Interface
			if arrV.Index(i).Interface() == elem {
				return true
			}
		}
	}

	return false
}

// AppTypeExists d
func AppTypeExists(appType string) bool {
	homeDir := GetWorkingDirNoError()

	appTypes, _ := ListDir(homeDir + "/template")
	for _, appT := range appTypes {
		if appT == appType {
			return true
		}
	}
	return false
}

// LibExists l
func LibExists(library string) bool {
	fmt.Println(string(os.PathSeparator))
	homeDir := GetWorkingDirNoError()

	libraryPath := filepath.Join(homeDir, "template", library, "codebase")

	if _, err := os.Stat(libraryPath); os.IsNotExist(err) {
		// path/to/whatever does not exist
		return false
	}
	return true

}

// ListDir will return list of directory inside path depth is 1
func ListDir(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var dirName []string
	for _, f := range files {
		if f.IsDir() {
			dirName = append(dirName, f.Name())
		}
	}
	return dirName, nil
}

// GetWorkingDir get current working directory
func GetWorkingDir() (string, error) {

	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path, nil
}

// GetWorkingDir get current working directory
func GetWorkingDirNoError() string {

	path, _ := os.Getwd()

	return path
}

// GetTemplateDir get root template directory
func GetTemplateDir() string {
	return filepath.Join(GetWorkingDirNoError(), consts.TemplatePath)
}

// AddCliLibs add these libs to supported cli libs
func AddCliLibs(cliNames []string) {
	consts.SupportedCliLib = cliNames
}

// GetOnlyTemplateCOnfigurableFiles this will list out files who will be configured in template
func GetOnlyTemplateCOnfigurableFiles(rootPath string) ([]string, error) {
	var filePath []string

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			//{{ this is the delimitor for template files
			//TODO: how to effieciently handle this process
			if strings.Contains(string(data), "{{") {
				filePath = append(filePath, path)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return filePath, nil
}
