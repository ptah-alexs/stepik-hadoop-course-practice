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
        strIn := strings.Split(scanner.Text(), "\t")
        if tempKey != strIn[0] {
            fmt.Println(strIn[0])
            tempKey = strIn[0]
        }
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
}
