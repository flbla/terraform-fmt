package handlers

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

func FormatCode(dir string) {
	varFiles := make(map[string][]*tfconfig.Variable)
	var vFiles []*variableFile
	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	module, err2 := tfconfig.LoadModule(dir)
	if err2 != nil {
		log.Println(err)
	}

	for _, v := range module.Variables {
		varFiles[v.Pos.Filename] = append(varFiles[v.Pos.Filename], v)
	}
	for file := range varFiles {
		vFiles = append(vFiles, newVariableFile(file, varFiles[file]))
	}

	for _, f := range vFiles {
		f.sortVariables()
	}
}
