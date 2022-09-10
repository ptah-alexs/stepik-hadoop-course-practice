package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
)

func main() {
    tempKey := ""
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        sVals := strings.Split(scanner.Text(), "\t")
        if tempKey != sVals[0] {
            fmt.Println(sVals[0])
            tempKey = sVals[0]
        }
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
}
