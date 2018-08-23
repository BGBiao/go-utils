/**
 * @File Name: iniconfigParse.go
 * @Author: xxbandy @http://xxbandy.github.io 
 * @Email:
 * @Create Date: 2018-08-23 18:08:26
 * @Last Modified: 2018-08-23 19:08:31
 * @Description:
 */

package main
import (
		"github.com/xxbandy/go-utils/confini"
    "fmt"
)

func main() {
        c, err := confini.NewFileConf("./conf.ini")
        if err != nil {
                fmt.Println(err)
                return
        }
        dbhost := c.String("db.host")
        dbname := c.String("db.name")
				dbpass := c.String("db.pass")

				port, err := c.Int("app.port")
        if err != nil {
                fmt.Println(err)
                return
        }

        applogdir := fmt.Sprintf("%s%s",c.String("app.logdir"),c.String("app.apppath"))
				fmt.Println(dbhost,dbname,dbpass,port,applogdir)
}
