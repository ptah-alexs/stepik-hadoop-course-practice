package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "strconv"
)

func main() {
    tempKey := ""
    sum := 0
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        v := strings.Split(scanner.Text(), "\t")
        key := v[0]
        value, _ := strconv.Atoi(v[1])
        if tempKey != "" && key != tempKey {
            fmt.Printf("%s\t%d\n", tempKey, sum)
            tempKey = key
            sum = value
        } else {
            sum += value
            tempKey = key
        }
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
    if tempKey != "" {
        fmt.Printf("%s\t%d\n", tempKey, sum)
    }
}
