package configuration

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type ApplicationProperties struct {
	Server struct {
		Port string
		Mode string
	}

	Database struct {
		Server       string
		Port         int
		User         string
		Password     string
		DatabaseName string `yaml:"database-name" `
		Pool         struct {
			MaxConnection     int `yaml:"max-connection"`
			MaxIdleConnection int `yaml:"max-idle-connection"`
		}
	}

	Redis struct {
		Host     string
		Port     int
		Password string
		DB       int
		Pool     struct {
			MaxIdle     int           `yaml:"max-idle"`
			IdleTimeout time.Duration `yaml:"idle-timeout"`
		}
	}

	Rabbit struct {
		Host     string
		Port     int
		Username string
		Password string
	}
}

var Properties *ApplicationProperties

func getProperties() *ApplicationProperties {
	properties := new(ApplicationProperties)
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

func init() {
	properties := getProperties()
	Properties = properties
}
