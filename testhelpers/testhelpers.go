package testhelpers

import (
	"io/ioutil"
	"os"
)

const (
	SAMPLE_PIPELINE = `---
jobs:
- name: run-postgres-task
  plan:
  - do:
    - task: ((var1))
      config:
        platform: linux
        image_resource:
          type: docker-image
          source:
            repository: postgres
        run:
          user: ((var2))
          path: sh
          args:
          - -exc
          - |
						whoami`

	SAMPLE_PIPELINE_MIXED = `---
jobs:
- name: run-postgres-task
  plan:
  - do:
    - task: {{var1}}
      config:
        platform: linux
        image_resource:
          type: docker-image
          source:
            repository: postgres
        run:
          user: ((var2))
          path: sh
          args:
          - -exc
          - |
						whoami`

	SAMPLE_PIPELINE_CURLYS = `---
jobs:
- name: run-postgres-task
  plan:
  - do:
    - task: {{var1}}
      config:
        platform: linux
        image_resource:
          type: docker-image
          source:
            repository: postgres
        run:
          user: {{var2}}
          path: sh
          args:
          - -exc
          - |
						whoami`

	SAMPLE_VARS = `var1: value1
var2: value2`

	SAMPLE_EXTRA_VARS = `var1: value1
var2: value2
var3: value3`

	SAMPLE_MISSING_VARS = `var1: value1`

	SAMPLE_VARS_PART1 = `var1: value1`

	SAMPLE_VARS_PART2 = `var2: value2`
)

func WriteStringToTempFile(folder, contents string) (string, error) {
	return WriteBytesToTempFile(folder, []byte(contents))
}
func WriteBytesToTempFile(folder string, contents []byte) (string, error) {
	tempFile, err := ioutil.TempFile(folder, "")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	err = ioutil.WriteFile(tempFile.Name(), contents, os.ModePerm)
	if err != nil {
		return "", err
	}

	return tempFile.Name(), err
}
