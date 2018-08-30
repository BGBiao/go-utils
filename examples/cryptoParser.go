/**
 * @File Name: cryptoParser.go
 * @Author: xxbandy @http://xxbandy.github.io 
 * @Email:
 * @Create Date: 2018-08-30 09:08:27
 * @Last Modified: 2018-08-30 09:08:03
 * @Description:
 */
package main
import (
		"github.com/xxbandy/go-utils/crypto"
    "fmt"
)


func main() {
	fmt.Println(crypto.Md5("hello"))
	fmt.Println(crypto.Hmac("key2", "hello"))
	fmt.Println(crypto.Sha1("hello"))
  fmt.Println(crypto.HmacSha1("key2","hello"))
}
