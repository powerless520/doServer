package unitTestingUtil

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http/httptest"
)

func HttpBuildQuery(param map[string]string) (url_str string) {
	value := ""
	for key, val := range param {
		value += "&" + key + "=" + val
	}
	tmp_url := value[1:]
	return tmp_url
}
func Get(uri string, router *gin.Engine) (body string, err error) {
	req := httptest.NewRequest("GET", uri, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result := w.Result()
	defer result.Body.Close()
	body_byte, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return "", err
	}
	return string(body_byte), nil
}
func PostForm(uri string, param map[string]string, router *gin.Engine) (body string, err error) {
	req := httptest.NewRequest("POST", uri+HttpBuildQuery(param), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result := w.Result()
	defer result.Body.Close()
	body_byte, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return "", err
	}
	return string(body_byte), nil
}
func PostBody(uri string, data string, router *gin.Engine) (body string, err error) {
	req := httptest.NewRequest("POST", uri, bytes.NewReader([]byte(data)))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result := w.Result()
	defer result.Body.Close()
	body_byte, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return "", err
	}
	return string(body_byte), nil
}
