package resources

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rcw5/pipeline-validator/utils"
)

type Pipeline struct {
	definition string
}

var SECRET_REGEXP = regexp.MustCompile(`(\(\([\w\.]+\)\)|{{[\w\.]+}})`)

func NewPipeline(definition string) Pipeline {
	return Pipeline{
		definition: definition,
	}
}

func (p Pipeline) Validate(vars ...Vars) utils.ValidationError {
	result := SECRET_REGEXP.FindAllStringSubmatch(p.definition, -1)

	var varsInPipeline []string
	for _, v := range result {
		varsInPipeline = append(varsInPipeline, strings.Trim(v[1], "({})"))
	}

	var allVars []string
	for _, s := range vars {
		allVars = append(allVars, s...)
	}
	undefinedVars, extraVars, _ := utils.CompareArrays(varsInPipeline, allVars)
	err := utils.ValidationError{}
	if undefinedVars != nil {
		err.MissingVarsError = fmt.Errorf("The following vars were present in the pipeline but not in the vars file: %s", strings.Join(undefinedVars, ", "))
	}
	if extraVars != nil {
		err.ExtraVarsError = fmt.Errorf("The following vars were present in the vars file but not the pipeline: %s", strings.Join(extraVars, ", "))
	}
	return err
}
