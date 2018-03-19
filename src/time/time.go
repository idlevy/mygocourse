package main

import (
    "fmt"
    "time"
)

func main() {

    t := time.Now()
    fmt.Printf("the time now is %s\n", t)
    fmt.Println(t.Month())
    fmt.Println(t.Day())


}
