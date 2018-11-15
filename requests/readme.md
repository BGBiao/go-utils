## 
`注意1:普通的请求可以直接使用net/http模块的get和post`

`注意2:特殊需求比如需要加header的需要构造特殊请求`


```
func main() {
    //baseurl := NewApi("http://dataapi.jd.com/v1/app?tenant=jrpe&username=xuxuebiao")
    baseurl := NewApi("http://dataapi.jd.com")
    token := `?tenant=jrpe&username=xuxuebiao`
    createreq,_ := baseurl.GetRequest("POST","/v1/app"+token,strings.NewReader("app=conda-repo"))
    createreq.Header.Add("Connection","keep-alive")
    respbody,_ := NewClient(createreq)
    fmt.Println(string(respbody))
}
```
