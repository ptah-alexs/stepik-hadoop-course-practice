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
        sVals := strings.Split(scanner.Text(), "\t")
        value, _ := strconv.Atoi(sVals[1])
        if tempKey != "" && sVals[0] != tempKey {
            parts := strings.Split(tempKey,"#")
            fmt.Printf("%s\t%s\t%d\n", parts[0], parts[1], sum)
            tempKey = sVals[0]
            sum = value
        } else {
            sum += value
            tempKey = sVals[0]
        }
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
    if tempKey != "" {
        parts := strings.Split(tempKey,"#")
        fmt.Printf("%s\t%s\t%d\n", parts[0], parts[1], sum)
    }
}
