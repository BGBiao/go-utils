/**
 * @File Name: test.go
 * @Author: xxbandy @http://xxbandy.github.io
 * @Email:
 * @Create Date: 2018-08-30 09:08:10
 * @Last Modified: 2018-08-30 09:08:22
 * @Description:
 */
package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func Md5(data string) string {
	md5 := md5.New()
	md5.Write([]byte(data))
	//md5Data := md5.Sum([]byte(""))
	//[]byte("") == nil
	md5Data := md5.Sum(nil)
	return hex.EncodeToString(md5Data)
}

func Hmac(key, data string) string {
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	return hex.EncodeToString(hmac.Sum(nil))
}

func Sha1(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum(nil))
}

func HmacSha1(key, data string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

