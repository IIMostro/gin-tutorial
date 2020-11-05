package configuration

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Properties struct {
	Server struct {
		Port string
		Mode string
	}

	Database struct {
		Type         string
		Server       string
		Name         string
		Password     string
		DatabaseName string `yaml:"database-name" `
	}
}

func GetProperties() *Properties {
	properties := new(Properties)
	file, err := ioutil.ReadFile("./configuration/application.yaml")
	if err != nil {
		log.Fatalf("read yaml error!, #%v", err)
	}
	err = yaml.Unmarshal(file, properties)
	if err != nil {
		log.Fatalf("unmarshal yaml error!, #%v", err)
	}

	return properties
}
