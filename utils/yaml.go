package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Xi float64 `yaml:"coefficient"`
}

func GetConfig() Config {
	// 读取YAML文件
	yamlFile, err := ioutil.ReadFile("../../config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// 解析YAML数据
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
