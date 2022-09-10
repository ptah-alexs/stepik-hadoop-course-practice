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
        if strings.Contains(scanner.Text(), "\tuser10\t") {
            fmt.Println(scanner.Text())
        }
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
}
