## 
`注意1:普通的请求可以直接使用net/http模块的get和post`

`注意2:特殊需求比如需要加header的需要构造特殊请求`

`特别注意:在发送http请求时，一定要合理的关闭tcp链接。一般在重定向时body也为nil，此时可能panic`

