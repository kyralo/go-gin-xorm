package utils

import (
	"io/ioutil"
	"log"
	"net/http"
)

/**
 * \* @author: WangChen
 * \* Date: 19-9-17
 * \* Time: 下午11:02
 */
const APPID  = ""
const APPSECRET = ""
const GEANTType  = ""

func GetBodyStr(jsCode string) []byte {

	response, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?" +
		"appid=" + APPID +
		"&secret=" + APPSECRET +
		"&js_code=" + jsCode +
		"&grant_type=" + GEANTType)



	if err != nil {
		log.Println(err)
		return  nil
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	return body
}