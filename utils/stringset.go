/**
 * @File Name: stringset.go
 * @Author: xxbandy @http://xxbandy.github.io 
 * @Email:
 * @Create Date: 2018-10-17 09:10:17
 * @Last Modified: 2018-10-17 10:10:34
 * @Description:
 */
package utils
import (
    //"fmt"
    //"os"
    //"reflect"
    "sort"
    "sync"
)


//golang实现hashset{https://blog.csdn.net/u012855229/article/details/52076784}
//golang中自定义set结构体以及相关的操作[hashset]
type StringSet struct {
    m map[string]bool
    sync.RWMutex
}

//初始化一个StringSet结构体指针
func NewStringSet() *StringSet {
    return &StringSet{
        m: map[string]bool{},
    }
}

//s 结构体拥有了sync.RWMutex读写锁,可以直接继承sync.RWMutex的相关方法
func (s *StringSet) Add(items ...string) {
    s.Lock()
    defer s.Unlock()
    if len(items) == 0 {
        return
    }
    for _, item := range items {
        s.m[item] = true
    }
}
func (s *StringSet) Remove(items ...string) {
    s.Lock()
    defer s.Unlock()
    if len(items) == 0 {
        return
    }
    for _, item := range items {
        delete(s.m, item)
    }
}

func (s *StringSet) Has(item string) bool {
    //读锁
    s.RLock()
    //读锁解锁
    defer s.RUnlock()
    _, ok := s.m[item]
    return ok
}

func (s *StringSet) Len() int {
    return len(s.List())
}

func (s *StringSet) Clear() {
    s.Lock()
    defer s.Unlock()
    s.m = map[string]bool{}
}

//判断结构体StringSet是否为空
func (s *StringSet) IsEmpty() bool {
    if s.Len() == 0 {
        return true
    }
    return false
}

//对结构体StringSet进行去重[set集合去重]
func (s *StringSet) List() []string {
    s.RLock()
    defer s.RUnlock()
    list := []string{}
    //s.m本身是一个有序的map[string]bool{}
    for item := range s.m {
        list = append(list, item)
    }
    return list
}

//去重加排序sort.Strings()对一个字符串数组进行排序[set集合去重加排序]
func (s *StringSet) SortList() []string {
    s.RLock()
    defer s.RUnlock()
    list := []string{}
    for item := range s.m {
        list = append(list, item)
    }
    sort.Strings(list)
    return list
}
