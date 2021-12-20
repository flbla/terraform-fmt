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

	for _, v := range module.Variables {
		fmt.Println(v.Name)
		fmt.Println(v.Description)
		fmt.Println(v.Pos.Filename)
		fmt.Println(v.Pos.Line)
	}
}
