/**
 * @File Name: sliceuniq.go
 * @Author: xxbandy @http://xxbandy.github.io 
 * @Email:
 * @Create Date: 2018-10-17 10:10:02
 * @Last Modified: 2018-10-17 11:10:31
 * @Description:
 */
package utils
import (
		"reflect"
)

//先排序后去重
func RemovalInt64(a []int64) (ret []int64) {
    //var ret  []int64{}
    for i := range a{
        flag := true
        for j := range ret {
            if a[i] == ret[j] {
                flag = false
                break
            }
        }
        if flag {
            ret = append(ret,a[i])
        }
    }
    return ret
}
func RemovalString(a []string) (ret []string) {
    for i := range a{
        flag := true
        for j := range ret {
            if a[i] == ret[j] {
                flag = false
                break
            }
        }
        if flag {
            ret = append(ret,a[i])
        }
    }
    return ret
}

//使用interface{}类型对slice进行排序[使用之之前需要把slice或list使用sort.Strings和sort.Ints进行转换]
func Duplicate(a interface{}) (ret []interface{}) {
   va := reflect.ValueOf(a)
   for i := 0; i < va.Len(); i++ {
      if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
         continue
      }
      ret = append(ret, va.Index(i).Interface())
   }
   return ret
}
