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
        skeys := strings.Split(scanner.Text(), "\t")
        for _, svals := range strings.Split(skeys[1], ",") {
            fmt.Printf("%s,%s\t1\n", strings.TrimSpace(skeys[0]), strings.TrimSpace(svals))
        }
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
}
