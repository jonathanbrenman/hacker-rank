package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(repeatedString("kmretasscityylpdhuwjirnqimlkcgxubxmsxpypgzxtenweirknjtasxtvxemtwxuarabssvqdnktqadhyktagjxoanknhgilnm", 736778906400))
}

func repeatedString(s string, n int64) int64 {
    if s == "a" {
        return n
    }
    counter := int64(0)
    for i := 0; i < len(s); i++ {
        if string(s[i]) == "a" {
            counter++
        }
    }
    predictTimes := float64(n) / float64(len(s))
    completeTimes, _ := math.Modf(predictTimes)
    counter *= int64(completeTimes)
    reminder := n % int64(len(s))
    for i := 0; i < int(reminder); i++ {
        if string(s[i]) == "a" {
            counter++
        }
    }
    return counter
}
