/**
 * @File Name: sliceuniq.go
 * @Author: xxbandy @http://xxbandy.github.io 
 * @Email:
 * @Create Date: 2018-10-17 10:10:02
 * @Last Modified: 2018-10-17 10:10:00
 * @Description:
 */
package utils

func RemovalSlice(a []string) (ret []string) {
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
