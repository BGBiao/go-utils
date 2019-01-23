package main

import (
	"fmt"
	"github.com/xxbandy/go-utils/utils"
	"sort"
)

func main() {
	list := []string{"xxbandy", "biaoge", "biao", "biaoge", "bgbiao"}
	if utils.Contains("biao", list) {
		fmt.Printf("%s is in %s\n", "biao", list)
	}
	//emovelist := utils.RemovalSlice(list)
	intlist := []int{1, 23, 13, 20, 10, 1, 23, 10}
	sort.Strings(list)
	sort.Ints(intlist)
	//removelist := utils.Duplicate(list)
	//removeintlist := utils.Duplicate(intlist)
	removelist := utils.RemovalString(list)
	removeintlist := utils.RemovalInt(intlist)
	fmt.Println(list, removelist, intlist, removeintlist)

	abc := utils.NewStringSet()
	abc.Add("xxbandy", "bgbiao", "biaoge", "bgbiao")
	setlist := abc.List()
	sortsetlist := abc.SortList()
	fmt.Println(setlist, sortsetlist)

  fmt.Printf("slice类型:%v %T",list,list)
  newlist,_ := utils.GetBytes(list)
  fmt.Printf("转换后的[]byte类型:%v %T",newlist,newlist)
}
