/**
 * @File Name: logger.go
 * @Author: xxbandy @http://xxbandy.github.io 
 * @Email:
 * @Create Date: 2018-08-24 06:08:45
 * @Last Modified: 2018-08-28 19:08:02
 * @Description:
 */
package logformat

import (
    "log"
    "os"
    "io"
)

/*
var (
    Trace   *log.Logger // 记录所有日志
    Info    *log.Logger // 重要的信息
    Warning *log.Logger // 需要注意的信息
    Error   *log.Logger // 致命错误
    Logfile string
)
*/
func LoggerInit(logfile,logLevel string) *log.Logger{
    file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalln("Failed to open error log file:", err)
    }

    //define the log output
    log.SetOutput(file)
    //define the log format 
    log.SetFlags(log.Ldate | log.Ltime |log.Lmicroseconds | log.Lshortfile)

    //log.SetPrefix("xxbandy")
    //os.Stdout could write a console and logfile.
    format := log.New(io.MultiWriter(file,os.Stdout), logLevel, log.Ldate | log.Ltime |log.Lmicroseconds | log.Lshortfile)
    return format 
}

