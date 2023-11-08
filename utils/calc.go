package utils

import (
	"math"
	"strings"
	"unicode/utf8"
)

/**
 * @author AronGe
 * @date 8/15/23  09:37
 * @description 获取相似度入口函数
 */
func GetSimilarity(s_new string, s_old string) float64 {
	var score float64 = 1
	if s_new == s_old {
		return score
	}
	return (calcScoreMain(s_old, s_new) + calcScoreMain(s_new, s_old)) / 2
}

/**
 * @author AronGe
 * @date 8/15/23  09:37
 * @description 获取相似度主计算函数
 */
func calcScoreMain(s_new string, s_old string) float64 {
	newLen := utf8.RuneCountInString(s_new)
	oldLen := utf8.RuneCountInString(s_old)
	newSLen := len(s_new)
	oldSLen := len(s_old)
	score := 0.0
	sOneNew := ""
	for i := 0; i < newLen; i++ {
		prefix := []rune(s_new)[i : i+1]
		// 将子串之前的字符串转换成[]rune
		sOneNew += string(prefix)
		strIndex := strings.Index(s_old, sOneNew)
		if strIndex == -1 {
			if utf8.RuneCountInString(sOneNew) >= 2 { //没找到且是连续2个字符以上的 截取前n-1个字符计算

				prefix1 := []rune(sOneNew)[0 : utf8.RuneCountInString(sOneNew)-1]
				// 将子串之前的字符串转换成[]rune
				contailStr := string(prefix1)
				score = calcScore(contailStr, score, strings.LastIndex(s_old, contailStr), strings.LastIndex(s_new, contailStr), newLen, oldLen, newSLen, oldSLen)
				i = i - 1
			}
			sOneNew = ""
		} else {
			if i == (newLen - 1) { //找到且是最后一个字符  直接拿来计算
				//fmt.Println(s_old,s_new, sOneNew,strings.Index(s_old,sOneNew), strings.Index(s_new,sOneNew) )
				score = calcScore(sOneNew, score, strings.LastIndex(s_old, sOneNew), strings.LastIndex(s_new, sOneNew), newLen, oldLen, newSLen, oldSLen)
			}
		}
	}
	return score
}

/**
 * @author AronGe
 * @date 8/15/23  09:37
 * @description 根据偏移量等参数获取单位相似度
 */

func calcScore(contailStr string, score float64, strIndexOld int, strIndexNew int, newLen int, oldLen int, newSLen int, oldSLen int) float64 {
	contailStrLen := utf8.RuneCountInString(contailStr)
	/*积分增加项= 长度匹配量 * 长度占比 越长越接近1*/
	asd := float64(contailStrLen) / (float64((newLen + oldLen)) / 2) * (float64(contailStrLen) / float64(contailStrLen+1))
	//fmt.Println(utils.Typeof(asd))
	//.Println(asd)
	score += asd
	pian := math.Abs(float64(strIndexOld - strIndexNew))
	//fmt.Println("偏移量",asd,pian)
	if pian != 0 {
		xi := pian / float64((newSLen+oldSLen)/2)
		asff := (asd * xi) / 6 //宽松度 越大越不敏感 越大分数越大  越不严格
		if asff < asd {
			score -= asff
			//fmt.Println(asff)
		}
	}
	//fmt.Println("/n")
	return score
}
