package main
import (
    "github.com/xxbandy/go-utils/utils"
    "fmt"
)

func main() {
    list := []string{"xxbandy","biaoge","biao","biaoge","bgbiao"}
    if utils.Contains("biao",list)   {
        fmt.Printf("%s is in %s\n","biao",list)
    }
    removelist := utils.RemovalSlice(list)
    fmt.Println(list,removelist)

    abc := utils.NewStringSet()
    abc.Add("xxbandy","bgbiao","biaoge","bgbiao")
    setlist := abc.List()
    sortsetlist := abc.SortList()
    fmt.Println(setlist,sortsetlist)
}

