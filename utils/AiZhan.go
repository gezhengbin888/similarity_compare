package utils

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

/*发送者邮箱*/
const AiKey = "d61285870222ae5ce4484a466d18e008"
const AiHost = "https://apistore.aizhan.com" //不要改此参数host

type AiData struct {
	BaiDuPc      float64 `json:"baidu_pc"`       //百度pc权重
	BaiDuMobile  float64 `json:"baidu_mobile"`   //百度移动权重
	SouGouPc     float64 `json:"sou_gou_pc"`     //搜狗pc权重
	SouGouMobile float64 `json:"sou_gou_mobile"` //搜狗移动权重
	San60Pc      float64 `json:"san_60_pc"`      //360pc权重
	San60Mobile  float64 `json:"san_60_mobile"`  //360移动权重
}

func GetHost() {

	// 指定要解压的目录
	directory := "/Users/gezhengbin/Desktop/360/src-vulnerability/tmp"

	// 递归遍历目录下的zip文件
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 检查文件是否为zip文件
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".zip") {
			// 解压zip文件
			if err := extractZipFile(path); err != nil {
				fmt.Printf("解压文件 %s 时出错: %v\n", path, err)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("遍历目录时出错: %v\n", err)
	}
}

func matchContent(content string, fName string) {
	// 定义用于匹配域名的正则表达式
	//regexPattern := `[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+`
	//regexPattern := `<span>(https?://)([^/]+)`
	regexPattern := `<span>(https?://)([^/:]+)[/:]`
	// 编译正则表达式
	regex, err := regexp.Compile(regexPattern)
	if err != nil {
		fmt.Println("正则表达式编译失败:", err)
		return
	}

	// 查找匹配的域名
	matches := regex.FindAllStringSubmatch(string(content), -1)
	// 提取域名
	if len(matches) >= 2 {
		domain := matches[1]
		if len(domain) >= 2 {
			GetBugAiBrs(domain[2], fName)
		}
	}
}

var i = 1

/*获取url的爱站权重值*/
func GetBugAiBrs(bugUrl string, fName string) AiData {
	var aiData AiData
	urlBaidu := AiHost + "/baidurank/siteinfos/" + AiKey + "?domains=" + bugUrl
	//urlSouGou := AiHost + "/sogourank/siteinfos/" + AiKey + "?domains=" + bugUrl
	//urlSan60 := AiHost + "/sorank/siteinfos/" + AiKey + "?domains=" + bugUrl
	aiData.BaiDuPc, aiData.BaiDuMobile = GetVrValue(urlBaidu)
	if aiData.BaiDuPc >= 1 || aiData.BaiDuMobile >= 1 {
		fmt.Println(fName, "--------------", bugUrl)
		// 构建目标文件的完整路径
		destinationFilePath := filepath.Join("/Users/gezhengbin/Desktop/360/src-vulnerability/ok/ok-now", filepath.Base(fName))

		// 移动文件
		err := moveFile(fName, destinationFilePath)
		if err != nil {
			fmt.Printf("移动文件失败：%v\n", err)
		}

	} else {
		fmt.Println(i)
	}
	//aiData.SouGouPc, aiData.SouGouMobile = GetVrValue(urlSouGou)
	//aiData.San60Pc, aiData.San60Mobile = GetVrValue(urlSan60)
	i++
	return aiData
}

/*根据url 获取爱站pc和移动权重值*/
func GetVrValue(url string) (float64, float64) {
	httpRes := HttpGet(url)
	code, ok1 := httpRes["code"].(float64)
	status, ok2 := httpRes["status"].(string)
	if httpRes != nil && ok1 && ok2 && code == 200000 && status == "success" {
		data1 := httpRes["data"].(map[string]interface{})
		data2 := data1["success"].([]interface{})
		if len(data2) >= 1 {
			data3 := data2[0].(map[string]interface{})
			return data3["pc_br"].(float64), data3["m_br"].(float64)
		}
	}
	return 0, 0
}

// HttpGet /*发送get请求*/
func HttpGet(path string) map[string]interface{} {
	res, err := http.Get(path)
	if err != nil {
		return StrToArrForRelation("")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return StrToArrForRelation("")
	}
	return StrToArrForRelation(string(body))
}

/*将字符串反解析为字典*/
func StrToArrForRelation(str string) map[string]interface{} {
	var d map[string]interface{}
	// 将字符串反解析为字典
	json.Unmarshal([]byte(str), &d)
	return d
}

func extractZipFile(zipFilePath string) error {
	r, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	defer r.Close()

	// 遍历zip文件中的文件
	for _, f := range r.File {
		if f.FileInfo().IsDir() {
			// 如果是文件夹，则递归解析第一个doc文件内容
			if err := extractFirstDocFile(f, zipFilePath); err != nil {
				return err
			}

			r, err := f.Open()
			if err != nil {
				return err
			}
			defer r.Close()

		} else {
			// 如果是文件，直接读取内容
			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return err
			}
			matchContent(string(data), zipFilePath)
		}
	}
	return nil
}

func extractFirstDocFile(folder *zip.File, zipFilePath string) error {
	// 在文件夹中查找第一个doc文件
	r, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	defer r.Close()

	if err != nil {
		return err
	}

	for _, file := range r.File {
		if strings.HasSuffix(file.Name, ".doc") {
			// 如果找到第一个doc文件，读取其内容
			rc, err := folder.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return err
			}
			matchContent(string(data), zipFilePath)
			break
		}
	}

	return nil
}

func moveFile(sourceFilePath, destinationFilePath string) error {
	// 打开源文件
	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// 创建目标文件
	destinationFile, err := os.Create(destinationFilePath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// 将源文件内容复制到目标文件
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	// 关闭源文件
	sourceFile.Close()

	// 删除源文件
	err = os.Remove(sourceFilePath)
	if err != nil {
		return err
	}

	return nil
}
