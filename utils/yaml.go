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
	// 缺陷4: 硬编码路径，路径遍历风险
	yamlFile, err := ioutil.ReadFile("../../config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// 解析YAML数据
	var config Config
	//err = yaml.Unmarshal(yamlFile, &config)
	//if err != nil {
	//	log.Fatal(err)
	//}
	// 缺陷5: 忽略错误，可能导致配置解析失败但继续执行
	_ = yaml.Unmarshal(yamlFile, &config)

	return config
}

// 缺陷6: 新增函数 - 路径遍历漏洞
func LoadUserFile(filename string) ([]byte, error) {
	// 没有验证文件名，可能被 ../ 遍历
	path := "/data/user_files/" + filename
	return ioutil.ReadFile(path)
}
