package handlers

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

func FormatCode(dir string) {
	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	module, err := tfconfig.LoadModule(dir)

	if err != nil {
		log.Println(err)
	}

	_ = sortVariables(module.Variables)
}
