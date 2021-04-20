package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func init() {
	unmarshal(configFile, &conf)
	unmarshal(linksFile, &r)
	unmarshal(headerFile, &h)
}

func unmarshal(path string, dest interface{}) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, dest)
	if err != nil {
		panic(err)
	}
}
