package handler

import (
	"fmt"
	"go-initializer/consts"
	"go-initializer/utils"
)

//RegisterCli to register all availabel clis
func RegisterCli() {

	consts.CliConfigurableFiles = make(map[string][]string)

	path, err := utils.GetWorkingDir()
	if err != nil {
		fmt.Println(err)
	}
	//getting all cliNames present in cliTemplatePath
	cliNames, err := utils.ListDir(path + consts.CliTempatePath)

	//add these libs to supported cli libs
	utils.AddCliLibs(cliNames)

	//reading files to get those file path which require templating
	for _, cliName := range cliNames {
		cliPath := fmt.Sprintf("%s%s/%s", path, consts.CliTempatePath, cliName)

		fileNames, err := utils.GetOnlyTemplateCOnfigurableFiles(cliPath)
		if err != nil {
			fmt.Println(err)
		}
		consts.CliConfigurableFiles[cliName] = fileNames
	}
}
