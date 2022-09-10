package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
)

func main() {
    tempKey := ""
    var dict map[string]int = map[string]int{}
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        strIn := strings.Split(scanner.Text(), "\t")
        ss := strings.TrimSpace(strIn[1])
        if tempKey != scanner.Text() {
            tempKey = scanner.Text()
            dict[ss]++
        }
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
    for key, val := range dict {
        fmt.Printf("%s\t%d\n", key, val)
    }
}
