/**
* @File Name: utils.go
* @Author: xxbandy @http://xxbandy.github.io
* @Email:
* @Create Date: 2018-11-14 12:11:22
* @Last Modified: 2018-11-15 22:11:27
* @Description:
type Response struct {
   Status     string // e.g. "200 OK"
   StatusCode int    // e.g. 200
   Proto      string // e.g. "HTTP/1.0"
   Header     Header
   Body       io.ReadCloser
   ContentLength int64
   Close bool
   Request *Request
}

type Request struct {
   Method string
   URL *url.URL
   Proto      string // "HTTP/1.0"
   Header Header
   Body io.ReadCloser
   GetBody func() (io.ReadCloser, error) // Go 1.8
   Host string
   Form url.Values
   PostForm url.Values
   MultipartForm *multipart.Form
   RemoteAddr string
   Response *Response
}

type Header map[string][]string

//package:url
type URL struct {
   Scheme     string
   Opaque     string
   User       *Userinfo
   Host       string
   Path       string
   RawPath    string
   ForceQuery bool
   RawQuery   string
   Fragment   string
}
*/
package requests

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	_ "strings"
	"time"
)

type reqApi struct {
	baseurl string
}

func NewApi(baseurl string) *reqApi {
	return &reqApi{
		baseurl: baseurl,
	}
}

//初始化特殊的http请求体
func (s *reqApi) GetRequest(method string, uri string, data io.Reader) (*http.Request, error) {
	//注意:由于后面的参数有?和&因此需要使用``方式避免字符串进行直接转义
	requrl := s.baseurl + uri
	req, reqErr := http.NewRequest(method, requrl, data)
	if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	return req, reqErr
}

//初始化http客户端
func NewClient(req *http.Request) ([]byte, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, respErr := client.Do(req)
	defer resp.Body.Close()
	if respErr != nil {
		return []byte{}, fmt.Errorf("%s", respErr)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

//普通的HTTP请求
//func Post(url, contentType string, body io.Reader) (resp *Response, err error)
//strings.NewReader("app=conda-repo")
func (s *reqApi) Post(uri, contenttype string, data io.Reader) ([]byte, error) {
	requrl := s.baseurl + uri
	resp, respErr := http.Post(requrl, contenttype, data)
	defer resp.Body.Close()
	if respErr != nil {
		return []byte{}, respErr
	}
	respdata, _ := ioutil.ReadAll(resp.Body)
	return respdata, nil
}

//func PostForm(url string, data url.Values) (resp *Response, err error)
func (s *reqApi) PostForm(uri, key, value string) ([]byte, error) {
	requrl := s.baseurl + uri
	postdata := url.Values{}
	postdata.Set(key, value)
	resp, respErr := http.PostForm(requrl, postdata)
	defer resp.Body.Close()
	if respErr != nil {
		return []byte{}, respErr
	}
	respdata, _ := ioutil.ReadAll(resp.Body)
	return respdata, nil

}

//func Get(url string) (resp *Response, err error)
func (s *reqApi) Get(uri string) ([]byte, error) {
	requrl := s.baseurl + uri
	resp, respErr := http.Get(requrl)
	defer resp.Body.Close()
	if respErr != nil {
		return []byte{}, respErr
	}
	respdata, _ := ioutil.ReadAll(resp.Body)
	return respdata, nil
}
