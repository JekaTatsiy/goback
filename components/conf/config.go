package conf

import (
	"io/ioutil"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func InitConfig(configFile string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.SetConfigFile(configFile)
	v.SetEnvPrefix("core")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// config := &Configuration{}

	// err = v.Unmarshal(config)
	// if err != nil {
	// 	return nil, err
	// }

	// return config, nil
	return v, nil
}

func NewConfig(file string, res interface{}) error {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, res)

	return err
}
