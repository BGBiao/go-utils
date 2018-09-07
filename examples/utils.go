package main
import (
    "github.com/xxbandy/go-utils/utils"
    "fmt"
)

func main() {
    list := []string{"xxbandy","biaoge","biao"}
    if utils.Contains("biao",list)   {
        fmt.Printf("%s is in %s\n","biao",list)
    }
}

