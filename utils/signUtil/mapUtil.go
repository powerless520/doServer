package signUtil

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func ObjToMapObj(paramObj interface{}) (data map[string]string) {
	jsonData, err := json.Marshal(paramObj)
	if err != nil {
		return
	}

	var obj map[string]interface{}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	decoder.Decode(&obj)

	data = mapToMapStr(obj)

	return
}

func mapToMapStr(paramMap map[string]interface{}) map[string]string {
	ret := make(map[string]string, len(paramMap))
	for k, v := range paramMap {
		if v != nil {
			ret[k] = fmt.Sprint(v)
		}
	}
	return ret
}
