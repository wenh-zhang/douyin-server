package pkg

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 md5加密
func MD5(src string) string {
	m := md5.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}