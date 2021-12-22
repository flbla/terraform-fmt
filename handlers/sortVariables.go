package handlers

import (
	"fmt"
	"sort"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

func sortVariables(variables map[string]*tfconfig.Variable) (err error) {
	files := make(map[string][]*tfconfig.Variable)
	names := make(map[string][]string)
	varMap := make(map[string]*tfconfig.Variable)

	for _, v := range variables {
		files[v.Pos.Filename] = append(files[v.Pos.Filename], v)
	}

	for k, _ := range files {
		for _, variable := range files[k] {
			names[k] = append(names[k], variable.Name)
			varMap[variable.Name] = variable
		}
		sort.Strings(names[k])
	}
	for k, _ := range files {
		for _, name := range names[k] {
			fmt.Println(varMap[name])
		}
	}

	return err
}
