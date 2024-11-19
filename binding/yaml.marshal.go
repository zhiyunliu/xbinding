package binding

import "gopkg.in/yaml.v3"

func (yamlBinding) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}
