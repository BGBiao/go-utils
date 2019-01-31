/**
 * @File Name: utils.go
 * @Author: xxbandy @http://xxbandy.github.io 
 * @Email:
 * @Create Date: 2019-01-31 12:01:18
 * @Last Modified: 2019-01-31 12:01:21
 * @Description:
 */

package utils
import (
		"time"
    "fmt"
    "encoding/gob"
    "os"
)


// format: 2006-01-02T 15:04:05
func ParserDatetoString(format string) string{
    now := time.Now().Format(format)
    return now
}

// append a content to a file
func WriteWithFileWrite(name, content string) {
    fileObj, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        fmt.Println("Failed to open the file", err.Error())
        LogError(err.Error(),name,"WriteWithFile")
        os.Exit(2)
    }
    defer fileObj.Close()
    if _, err := fileObj.WriteString(content); err == nil {
        return
    }
    return
}


