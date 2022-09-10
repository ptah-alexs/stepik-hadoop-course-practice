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
    count := 0
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        v := strings.Split(scanner.Text(), "\t")
        key := v[0]
        value, err := strconv.Atoi(v[1])
        if err != nil {
            panic(err)
        }
        if tempKey != "" && key != tempKey {
            fmt.Printf("%s\t%d\n", tempKey, sum/count)
            tempKey = key
            sum = value
            count = 1
        } else {
            sum += value
            tempKey = key
            count++
        }
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
    if tempKey != "" {
        fmt.Printf("%s\t%d\n", tempKey, sum/count)
    }
}
