package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        sVals := strings.Split(scanner.Text(), "\t")
        fmt.Println(strings.TrimSpace(sVals[2]))
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
}
