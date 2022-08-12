package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 未加盐的MD5
func Md5(v string) string {
	data := md5.Sum([]byte(v))
	return hex.EncodeToString(data[:])
}

// 加盐的MD5
func Md5Salt(v, salt string) string {
	hash := md5.New()
	hash.Write([]byte(v))
	hash.Write([]byte(salt))
	return hex.EncodeToString(hash.Sum(nil))
}
