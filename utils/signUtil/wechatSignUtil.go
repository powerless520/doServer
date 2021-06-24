package signUtil

import (
	"reflect"
	"sort"
	"strings"
)

func Struct2MapString(obj interface{}) map[string]string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		data[snakeString(t.Field(i).Name)] = v.Field(i).String()
	}
	return data
}

// WechatNotifySignCreate ...
func WechatNotifySignCreate(obj interface{}, keyStr string) (signStr, signNew string) {
	var params map[string]string
	if reflect.ValueOf(obj).Kind() != reflect.Map {
		params = Struct2MapString(obj)
	} else {
		params = obj.(map[string]string)
	}
	var keySlice []string
	for k := range params {
		keySlice = append(keySlice, k)
	}
	sort.Strings(keySlice)
	signStr = ""
	for _, k := range keySlice {
		if k != "sign" && params[k] != "" {
			signStr += "&" + k + "=" + params[k]
		}
	}
	signStr = strings.TrimLeft(signStr, "&")
	signStr += "&key=" + keyStr
	signNew = strings.ToUpper(Md5Encode(signStr))
	return
}
func KsortStrHttpBuildQuery(params map[string]string) (str string) {
	var key_slice []string
	for k, _ := range params {
		key_slice = append(key_slice, k)
	}
	sort.Strings(key_slice)
	for _, k := range key_slice {
		if k != "sign" {
			str += "&" + k + "=" + params[k]
		}
	}
	str = strings.TrimLeft(str, "&")
	return
}
func KsortAndImplodeVal(params map[string]string) (str string) {
	var key_slice []string
	for k, _ := range params {
		key_slice = append(key_slice, k)
	}
	sort.Strings(key_slice)
	for _, k := range key_slice {
		if k != "sign" {
			str += params[k]
		}
	}
	return
}
func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}
