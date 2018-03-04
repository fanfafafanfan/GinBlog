package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router := gin.Default()
	router.GET("/", IndexAPI)
	router.POST("/article/add", AddarticleAPI)
	router.POST("/article/searchAll", GetarticlesAPI)
	router.POST("/article/searchById", GetarticleAPI)
	router.POST("/article/update", ModarticleAPI)
	router.POST("/article/delete", DelarticleAPI)
}
func ParseToStr(mp map[string]string) string {
	values := ""
	for key, val := range mp {
		values += "&" + key + "=" + val
	}

	temp := values[1:]
	values = "?" + temp
	return values
}
func FormRequest(method string, uri string, param map[string]string, router *gin.Engine) []byte {
	req := httptest.NewRequest(method, uri+ParseToStr(param), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result := w.Result()
	defer result.Body.Close()
	body, _ := ioutil.ReadAll(result.Body)
	return body
}

type Req struct {
	message string `json:"msg"`
}

func TestAddarticleAPI(t *testing.T) {
	url := "/article/add"
	param := make(map[string]string)
	param["ArticleName"] = "Package logging"
	param["ArticleContent"] = "Package logging implements a logging infrastructure for Go. "
	param["Author"] = "Lily"
	body := FormRequest("POST", url, param, router)
	fmt.Printf("response:%v\n", string(body))
	str := "新增成功!"
	r := &Req{}
	err := json.Unmarshal(body, r)
	if r.message != str {
		t.Errorf(" 响应字符串不符，body:%v\n", string(body))
	}
	if err != nil {
		log.Fatalln(err)
	}
}
func TestGetarticlesAPI(t *testing.T) {
	url := "/article/searchAll"
	param := make(map[string]string)
	body := FormRequest("POST", url, param, router)
	fmt.Printf("response:%v\n", string(body))
	r := &Req{}
	err := json.Unmarshal(body, r)
	if r.message != "success" {
		t.Errorf(" 响应字符串不符，body:%v\n", string(body))
	}
	if err != nil {
		log.Fatalln(err)
	}
}
func TestGetarticleAPI(t *testing.T) {
	url := "/article/searchById"
	param := make(map[string]string)
	param["Id"] = "1"
	body := FormRequest("POST", url, param, router)
	fmt.Printf("response:%v\n", string(body))
	r := &Req{}
	err := json.Unmarshal(body, r)
	if r.message != "success" {
		t.Errorf(" 响应字符串不符，body:%v\n", string(body))
	}
	if err != nil {
		log.Fatalln(err)
	}
}
func TestModarticleAPI(t *testing.T) {
	url := "/article/update"
	param := make(map[string]string)
	param["Id"] = "1"
	body := FormRequest("POST", url, param, router)
	fmt.Printf("response:%v\n", string(body))
	str := "更新成功!"
	r := &Req{}
	err := json.Unmarshal(body, r)
	if r.message != str {
		t.Errorf(" 响应字符串不符，body:%v\n", string(body))
	}
	if err != nil {
		log.Fatalln(err)
	}
}
func TestDelarticleAPI(t *testing.T) {
	url := "/article/delete"
	param := make(map[string]string)
	param["Id"] = "2"
	body := FormRequest("POST", url, param, router)
	fmt.Printf("response:%v\n", string(body))
	str := "删除成功!"
	r := &Req{}
	err := json.Unmarshal(body, r)
	if r.message != str {
		t.Errorf(" 响应字符串不符，body:%v\n", string(body))
	}
	if err != nil {
		log.Fatalln(err)
	}
}
