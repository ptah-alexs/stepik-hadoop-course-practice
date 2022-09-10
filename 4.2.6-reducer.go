package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
)

func main() {
    tempKey := ""
    var dict map[string]string = map[string]string{}
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        sVals := strings.Split(scanner.Text(), "\t")
        if tempKey != sVals[0] {
            if dict[tempKey] == "A" {
                fmt.Println(tempKey)
            }
            tempKey = sVals[0]
            dict = map[string]string{}
        }
        dict[sVals[0]]+=sVals[1]
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
    if tempKey != "" {
        if dict[tempKey] == "A" {
            fmt.Println(tempKey)
        }
    }
}
