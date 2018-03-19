package main

import(
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().Unix())
    dow := rand.Intn(6)+1
    fmt.Println("Day", dow)


    result := ""
    switch dow {
        case 1:
            result="its a sunday"
        case 7:
            result="its a suterday"
        default:
            result = "its a weekday :("
    }
    fmt.Println("Day", dow ,",", result)
}
