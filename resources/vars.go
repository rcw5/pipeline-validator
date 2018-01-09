package resources

import yaml "gopkg.in/yaml.v2"

type Vars []string

func NewVars(secretsYAML string) (Vars, error) {
	var vars map[string]interface{}
	err := yaml.Unmarshal([]byte(secretsYAML), &vars)
	if err != nil {
		return Vars{}, err
	}
	var keys []string
	for k := range vars {
		keys = append(keys, k)
	}

	return keys, err
}
