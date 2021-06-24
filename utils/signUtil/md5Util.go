package signUtil

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Encode(msg string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(msg))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
