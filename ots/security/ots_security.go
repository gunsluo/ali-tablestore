package security

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
)

func ContentMd5(buffer []byte) string {
	md5 := md5.Sum(buffer)
	return base64.StdEncoding.EncodeToString(md5[:])
}

func HmacSha1(key, message []byte) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}

// 签名
func Signature(accessKeySecret, stringToSign string) string {
	mac := HmacSha1([]byte(accessKeySecret), []byte(stringToSign))
	return base64.StdEncoding.EncodeToString(mac)
}
