/**
 * @File Name: logger.go
 * @Author: xxbandy @http://xxbandy.github.io 
 * @Email:
 * @Create Date: 2018-08-24 06:08:35
 * @Last Modified: 2018-08-28 19:08:59
 * @Description:
 */
package main
import (
    "github.com/xxbandy/go-utils/logformat"
)

func Logger(level,format string,message ...interface{}) {
    logger := logformat.LoggerInit("detail.log",level)
    logger.Printf(format,message...)

}

func main() {
    Logger("Info: ","somethings:%s occured!\n%s","hello","github.com/xxbandy/go-utils/logformat") 
    }
