/**
 * @File Name: request.go
 * @Author: xxbandy @http://xxbandy.github.io
 * @Email:
 * @Create Date: 2018-11-15 18:11:15
 * @Last Modified: 2018-11-15 23:11:12
 * @Description:
 */
package main

import (
	"fmt"
	"github.com/xxbandy/go-utils/requests"
	_ "strings"
)

func main() {
	/*
	   //post相关方法
	   baseurl := requests.NewApi("http://data.xxbandy.com")
	   token := `?tenant=jrpe&username=bgbiao`
	   uri := "/v1/app"+token

	   //1. post with header
	   createreq,_ := baseurl.GetRequest("POST",uri,strings.NewReader("app=conda-repo"))
	   createreq.Header.Add("Connection","keep-alive")
	   respbody,_ := requests.NewClient(createreq)
	   fmt.Println(string(respbody))

	   //2. base post
	   data,_ := baseurl.Post(uri,"application/x-www-form-urlencoded",strings.NewReader("app=conda-repo"))
	   fmt.Println(string(data))

	   //3. postform
	   data,_ := baseurl.PostForm(uri,"app","conda-repo")
	   fmt.Println(string(data))

	*/

	/*
	   //get相关方法
	*/
	baseurl := requests.NewApi("http://www.cip.cc")

	//1. get with header
	createreq, _ := baseurl.GetRequest("GET", "/", nil)
	createreq.Header.Add("Connection", "keep-alive")
	//注意:如果增加UA,会根据UA判断返回那种数据格式，比如curl会返回接口类型数据,不加这个UA，会返回html
	createreq.Header.Add("User-Agent", "curl")
	data, _ := requests.NewClient(createreq)
	fmt.Println(string(data))

	//2. base get
	//data,_ := baseurl.Get("/")
	//返回的内容为整个html内容
	//fmt.Println(string(data))

}
