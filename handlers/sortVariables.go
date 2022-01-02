package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/zclconf/go-cty/cty"
	ctyjson "github.com/zclconf/go-cty/cty/json"
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

func getCtyVal(inputValue interface{}) (defaultValue cty.Value, err error) {
	if inputValue == nil {
		return defaultValue, err
	}

	defaultValueJson, err := json.Marshal(inputValue)
	if err != nil {
		return defaultValue, err
	}

	implType, err := ctyjson.ImpliedType(defaultValueJson)
	if err != nil {
		return defaultValue, err
	}

	defaultValue, err = ctyjson.Unmarshal(defaultValueJson, implType)
	if err != nil {
		return defaultValue, err
	}
	return defaultValue, err
}

func (f *variableFile) sortVariables() (err error) {
	var names []string
	varMap := make(map[string]*tfconfig.Variable)

	for _, variable := range f.variable {
		names = append(names, variable.Name)
		fmt.Println(variable.Default)
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
		defaultValue, err := getCtyVal(varMap[name].Default)
		if err != nil {
			return err
		}
		b.Body().SetAttributeValue("default", defaultValue)
		for _, v := range varMap[name].Validations {
			val, err := getCtyVal(v)

			if err != nil {
				return err
			}
			b.Body().SetAttributeValue("validation", val)
		}
		b.Body().SetAttributeValue("sensitive", cty.BoolVal(varMap[name].Sensitive))
		file.Body().AppendBlock(b)
		file.Body().AppendNewline()
	}
	outputFile.Write(file.Bytes())

	return err
}
