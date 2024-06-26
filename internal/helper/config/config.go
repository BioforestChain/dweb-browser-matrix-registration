package config

import (
	"fmt"
	"github.com/BioforestChain/dweb-browser-matrix-service-registration/pkg/support-go/bootstrap"
	"github.com/BioforestChain/dweb-browser-matrix-service-registration/pkg/support-go/helper/config"
	"io/ioutil"
	"os"
)

func GetDomainConfig(fileName string) []string {
	list := make([]string, 0)
	path := fmt.Sprintf("%sconfig/%s/%s.yml", bootstrap.ProjectPath(), bootstrap.DevEnv, fileName)
	DBConfigs, err := config.GetConfig(path)
	if err != nil {
		return list
	}
	configList, err := DBConfigs.List("list")
	if err != nil {
		return list
	}
	for _, v := range configList {
		list = append(list, v.(string))
	}

	return list

}

func GetDomainMark(fileName string) string {
	var domainMark string
	path := fmt.Sprintf("%sconfig/%s/%s.yml", bootstrap.ProjectPath(), bootstrap.DevEnv, fileName)
	Configs, err := config.GetConfig(path)
	if err != nil {
		return domainMark
	}
	domainMark, _ = Configs.String("domainMark")
	return domainMark

}
func GetKeyValue(fileName string, keyName string) string {
	var domainMark string
	path := fmt.Sprintf("%sconfig/%s/%s.yml", bootstrap.ProjectPath(), bootstrap.DevEnv, fileName)
	Configs, err := config.GetConfig(path)
	if err != nil {
		return domainMark
	}
	domainMark, _ = Configs.String(keyName)
	return domainMark
}

func GetContentFromPem(fileName string) string {
	path := fmt.Sprintf("%sconfig/%s/%s.pem", bootstrap.ProjectPath(), "certs", fileName)
	// 获取文件信息
	file, err := os.Stat(path)
	if err != nil {
		fmt.Println("file err :", err)
		return ""
	}
	// 判断是否是文件
	if !file.IsDir() {
		// 读取文件内容
		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println("file content err :", err)
			return ""
		}
		return string(content)
	}
	return ""
}
