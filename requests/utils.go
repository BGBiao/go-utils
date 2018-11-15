/**
 * @File Name: utils.go
 * @Author: xxbandy @http://xxbandy.github.io 
 * @Email:
 * @Create Date: 2018-11-14 12:11:22
 * @Last Modified: 2018-11-15 18:11:05
 * @Description:
 */
package requests
import (
		"net/http"
		"net/url"
    "fmt"
    "io"
    "time"
    "io/ioutil"
    "strings"

)


type reqApi struct {
    baseurl         string
}

func NewApi(baseurl string) *reqApi {
    return &reqApi{
        baseurl: baseurl,
    }
}


//初始化特殊的http请求体
func (s *reqApi) GetRequest(method string,uri string,data io.Reader) (*http.Request, error){
    //http://hostname /v1/app ?tenant=jrpe&username=xuxuebiao
    //注意:由于后面的参数有?和&因此需要使用``方式避免字符串进行直接转义
    requrl := s.baseurl+uri
    req,reqErr := http.NewRequest(method,requrl,data)
    if method == "POST" {
        req.Header.Set("Content-Type","application/x-www-form-urlencoded")
    }
		return req,reqErr
}

//初始化http客户端
func NewClient(req *http.Request) ([]byte,error) {
		client := &http.Client{
				Timeout:		5 * time.Second,
		}
		resp,respErr := client.Do(req)
		defer resp.Body.Close()
		if respErr != nil{
				return []byte{},fmt.Errorf("%s",respErr)
		}
	  body,_ := ioutil.ReadAll(resp.Body)	
		return body,nil
}

//普通的HTTP请求
//func Post(url, contentType string, body io.Reader) (resp *Response, err error)
//strings.NewReader("app=conda-repo")
func (s *reqApi) Post(uri,contenttype string,data io.Reader) ([]byte,error) {
    requrl := s.baseurl+uri
    resp,respErr := http.Post(requrl,contenttype,data)
    defer resp.Body.Close()
    if respErr != nil {
        return []byte{},respErr
    }
    respdata,_ := ioutil.ReadAll(resp.Body)
    return respdata,nil
}

//func PostForm(url string, data url.Values) (resp *Response, err error)
func (s *reqApi) PostForm(uri,key,value string) ([]byte,error) {
    requrl := s.baseurl+uri
    postdata := url.Values{}
    postdata.Set(key,value)
    resp,respErr := http.PostForm(requrl, postdata)
    defer resp.Body.Close()
    if respErr != nil {
        return []byte{},respErr
    }
    respdata,_ := ioutil.ReadAll(resp.Body)
		return respdata,nil

}

//func Get(url string) (resp *Response, err error)
func (s *reqApi) Get(uri string) ([]byte,error) {
		requrl := s.baseurl+uri
		resp,respErr := http.Get(requrl)
		defer resp.Body.Close()
    if respErr != nil {
        return []byte{},respErr
    }
    respdata,_ := ioutil.ReadAll(resp.Body)
    return respdata,nil
}



/*
func main() {
    //baseurl := NewApi("http://dataapi.jd.com/v1/app?tenant=jrpe&username=xuxuebiao")
    baseurl := NewApi("http://dataapi.jd.com")
		token := `?tenant=jrpe&username=xuxuebiao`
    createreq,_ := baseurl.GetRequest("POST","/v1/app"+token,strings.NewReader("app=conda-repo"))
    createreq.Header.Add("Connection","keep-alive")
    respbody,_ := NewClient(createreq)
    fmt.Println(string(respbody))
}

*/
