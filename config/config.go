package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	CorpId              string `yaml:"CorpId"`
	ProviderSecret      string `yaml:"ProviderSecret"`
	SuiteId             string `yaml:"SuiteId"`
	SuiteSecret         string `yaml:"SuiteSecret"`
	SuiteToken          string `yaml:"SuiteToken"`
	SuiteEncodingAesKey string `yaml:"SuiteEncodingAesKey"`
	Dsn                 string `yaml:"Dsn"`
	Port                int    `yaml:"Port"`
}

func ParseFile(yml string) (c *Config) {
	if yml == "" {
		yml = "suite.yml"
	}
	yamlFile, err := ioutil.ReadFile(yml)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		panic(err)
	}
	return c
}
