package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func NewConfig(file string, res interface{}) error {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, res)

	return err
}
