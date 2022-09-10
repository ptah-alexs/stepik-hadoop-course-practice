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
        v := strings.Split(scanner.Text(), " ")
        for _, w := range v {
            if w != "" {
                fmt.Printf("%s\t1\n", strings.TrimSpace(w))
            }
        }
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
}
