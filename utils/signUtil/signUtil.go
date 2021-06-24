package signUtil

import (
	"sort"
	"strings"
)

func GetSignStr(data map[string]string, includeEmptyParam bool, joinSep string) string {
	sortedKeys := Ksort(data)
	values := make([]string, 0)

	for _, v := range sortedKeys {
		if includeEmptyParam || data[v] != "" {
			values = append(values, v+"="+data[v])
		}
	}

	signStr := strings.Join(values, joinSep)
	return signStr
}

func GetSignStrWithoutKey(data map[string]string, includeEmptyParam bool, joinSep string) string {
	sortedKeys := Ksort(data)
	values := make([]string, 0)

	for _, v := range sortedKeys {
		if includeEmptyParam || data[v] != "" {
			values = append(values, data[v])
		}
	}

	signStr := strings.Join(values, joinSep)
	return signStr
}

func XmlCreateSignStr(obj map[string]string, keyStr string) (signStr, signNew string) {
	var params map[string]string

	var keySlice []string
	for k := range obj {
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
