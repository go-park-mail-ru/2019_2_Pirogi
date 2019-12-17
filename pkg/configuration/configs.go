package configuration

import (
	"io/ioutil"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"gopkg.in/yaml.v2"
)

func UnmarshalConfigs(configsPath string) error {
	if string(configsPath[len(configsPath)-1]) != "/" {
		configsPath += "/"
	}
	file, err := ioutil.ReadFile(configsPath + "default.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(file, &configs.Default)
	if err != nil {
		return err
	}

	file, err = ioutil.ReadFile(configsPath + "headers.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(file, &configs.Headers)
	if err != nil {
		return err
	}
	return nil
}
