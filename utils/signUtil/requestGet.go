package signUtil

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
)

func ReqSignGet(params map[string]string) (sign_str, sign_new string) {
	var err error
	params_length := len(params)
	val_ary := make([]string, params_length)
	var i int = 0
	for k, v := range params {
		if k == "app_callback_url" {
			v, err = url.QueryUnescape(v)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		if v != "" && k != "sign" && k != "key" && k != "paid" && k != "action" && k != "resource_id" && k != "extra_currency" && k != "cash_type" && k != "callback_url" && k != "jump_url" && k != "app_name" && k != "app_user_name" && k != "product_name" && k != "user_ip" && k != "userip" && k != "xyzs_order_time" && k != "xyzs_deviceid" {
			val_ary[i] = v
		}
		i++
	}

	sort.Strings(val_ary)
	sign_str = "582df15de91b3f12d8e710073e43f4f8" + strings.Join(val_ary, "")
	sign_new = Md5Encode(sign_str)
	return
}
func RequestSignGet(params map[string]string) (sign string) {
	var err error
	params_length := len(params)
	val_ary := make([]string, params_length)
	var i int = 0
	for k, v := range params {
		if k == "app_callback_url" {
			v, err = url.QueryUnescape(v)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		if v != "" && k != "sign" && k != "key" && k != "paid" && k != "action" && k != "resource_id" && k != "extra_currency" &&
			k != "cash_type" && k != "callback_url" && k != "jump_url" && k != "app_name" && k != "app_user_name" &&
			k != "product_name" && k != "user_ip" && k != "userip" && k != "xyzs_order_time" && k != "xyzs_deviceid" {
			val_ary[i] = v
		}
		i++
	}

	sort.Strings(val_ary)
	sign_str := strings.Join(val_ary, "")
	sign = Md5Encode("582df15de91b3f12d8e710073e43f4f8" + sign_str)
	return sign
}
