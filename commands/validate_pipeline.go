package commands

import (
	"errors"
	"io/ioutil"

	"github.com/rcw5/pipeline-validator/resources"
)

func ValidatePipeline(pipelinePath string, varsFilePaths ...string) error {
	allVars := []resources.Vars{}
	for _, varsFile := range varsFilePaths {
		contents, err := ioutil.ReadFile(varsFile)
		if err != nil {
			return err
		}
		vars, err := resources.NewVars(string(contents))
		if err != nil {
			return err
		}
		allVars = append(allVars, vars)
	}

	contents, err := ioutil.ReadFile(pipelinePath)
	if err != nil {
		return err
	}
	pipelineObj := resources.NewPipeline(string(contents))
	validationError := pipelineObj.Validate(allVars...)
	if !validationError.IsSuccessful() {
		return errors.New(validationError.Error())
	}
	return nil
}
