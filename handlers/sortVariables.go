package handlers

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/zclconf/go-cty/cty"
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
	file := hclwrite.NewFile()

	for k, _ := range files {
		f, err := os.Create(k + ".sorted")

		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		for _, name := range names[k] {
			b := hclwrite.NewBlock("variable", []string{varMap[name].Name})
			b.Body().SetAttributeValue("type", cty.StringVal(varMap[name].Type))
			b.Body().SetAttributeValue("description", cty.StringVal(varMap[name].Description))
			b.Body().SetAttributeValue("default", cty.StringVal(fmt.Sprint(varMap[name].Default)))
			b.Body().SetAttributeValue("sensitive", cty.BoolVal(varMap[name].Sensitive))
			file.Body().AppendBlock(b)
			file.Body().AppendNewline()
		}
		f.Write(file.Bytes())
	}

	return err

}
