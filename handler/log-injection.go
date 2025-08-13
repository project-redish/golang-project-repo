package handler

import (
	"encoding/json"
	"fmt"
	"go-initializer/types"
	"go-initializer/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)

func (request *GenerateTemplateRequest) appendLogGoFile() error {

	goFileContent, err := getLogGoFile(request.LoggingFramework)

	if err != nil {
		return fmt.Errorf("Error getting log file content for framework %s with error %s", request.LoggingFramework, err.Error())
	}

	if _, err := os.Stat(filepath.Join(request.outputFolder, "logger")); os.IsNotExist(err) {
		err = os.Mkdir(filepath.Join(request.outputFolder, "logger"), os.ModePerm)
		if err != nil {
			return fmt.Errorf("Cannot create logging directory for request %+v with error %s", request, err.Error())
		}
	}

	file, err := os.Create(filepath.Join(request.outputFolder, "logger/logger.go"))
	if err != nil {
		return fmt.Errorf("Cannot create logging file inside output folder for request %+v with error %s", request, err.Error())
	}
	defer file.Close()
	_, err = file.Write(goFileContent)
	if err != nil {
		return fmt.Errorf("Error writing contents to output logging file for request %+v with error %s", request, err.Error())
	}
	return nil
}

func getLogGoFile(logFilePath string) ([]byte, error) {

	homePath := utils.GetWorkingDirNoError()

	goFilePath := filepath.Join(homePath, "template", "logframework", logFilePath, "logger.go")

	return ioutil.ReadFile(goFilePath)
}

func readLogJson(logFramework string) (*types.LoggingFramework, error) {

	homePath := utils.GetWorkingDirNoError()
	jsonFilePath := filepath.Join(homePath, "template", "logframework", logFramework, "logger.json")

	types.Mutex.Lock()
	
	jsonData, err := ioutil.ReadFile(jsonFilePath)
	
	types.Mutex.Unlock()
	
	if err != nil {
		return nil, fmt.Errorf("Error while reading json log file for logging framework %s with error %s", logFramework, err.Error())
	}

	var frameworkData types.LoggingFramework

	err = json.Unmarshal([]byte(jsonData), &frameworkData)

	if err != nil {
		return nil, fmt.Errorf("Error Unmarshalling  json log file for logging framework %s with error %s", logFramework, err.Error())
	}

	return &frameworkData, nil

}
