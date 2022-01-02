package handlers

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/zclconf/go-cty/cty"
)

type variableFile struct {
	file     string
	variable []*tfconfig.Variable
}

func newVariableFile(file string, variable []*tfconfig.Variable) *variableFile {
	return &variableFile{
		file:     file,
		variable: variable,
	}
}

func (f *variableFile) sortVariables() (err error) {
	var names []string
	varMap := make(map[string]*tfconfig.Variable)

	for _, variable := range f.variable {
		names = append(names, variable.Name)
		varMap[variable.Name] = variable
	}
	sort.Strings(names)

	file := hclwrite.NewFile()
	outputFile, err := os.Create(f.file + ".sorted")

	if err != nil {
		log.Fatal(err)
	}

	defer outputFile.Close()
	for _, name := range names {
		b := hclwrite.NewBlock("variable", []string{varMap[name].Name})
		b.Body().AppendUnstructuredTokens(hclwrite.Tokens{{
			Type:  hclsyntax.TokenNil,
			Bytes: []byte(fmt.Sprintf("type = %s", varMap[name].Type)),
		}})
		b.Body().AppendNewline()
		b.Body().SetAttributeValue("description", cty.StringVal(varMap[name].Description))
		if strings.HasPrefix(varMap[name].Type, "object") {
			//v, _ := json.Marshal(varMap[name].Default)
			b.Body().SetAttributeRaw("default", hclwrite.Tokens{{
				Type:  hclsyntax.TokenNil,
				Bytes: []byte(fmt.Sprint(varMap[name].Default)),
			}})

			b.Body().AppendNewline()
		} else {
			b.Body().SetAttributeValue("default", cty.StringVal(fmt.Sprint(varMap[name].Default)))
		}
		b.Body().SetAttributeValue("sensitive", cty.BoolVal(varMap[name].Sensitive))
		file.Body().AppendBlock(b)
		file.Body().AppendNewline()
	}
	outputFile.Write(file.Bytes())

	return err
}
