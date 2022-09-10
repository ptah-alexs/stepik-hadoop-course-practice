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
        sKeys := strings.Split(scanner.Text(), "\t")
        sVals := strings.Split(sKeys[0], ",")
        fmt.Printf("%s\t1\n", strings.TrimSpace(sVals[1]),)
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
}
