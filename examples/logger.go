/**
 * @File Name: logger.go
 * @Author: xxbandy @http://xxbandy.github.io 
 * @Email:
 * @Create Date: 2018-08-24 06:08:35
 * @Last Modified: 2018-08-24 06:08:43
 * @Description:
 */
package main
import (
    "github.com/xxbandy/go-utils/logformat"
)

func main() {
    logger := logformat.LoggerInit("detail.log","Info: ")
    logger.Println("hello")
    
    }
