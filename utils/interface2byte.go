/**
 * @File Name: interface2byte.go
 * @Author: xxbandy @http://xxbandy.github.io 
 * @Email:
 * @Create Date: 2019-01-23 15:01:41
 * @Last Modified: 2019-01-23 15:01:12
 * @Description:
 */
package utils
import (
    "encoding/gob"
    "bytes"
)

func GetBytes(key interface{}) ([]byte, error) {
    var buf bytes.Buffer
    enc := gob.NewEncoder(&buf)
    err := enc.Encode(key)
    if err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}
