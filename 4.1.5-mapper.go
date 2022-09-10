package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
)

func main() {
    var dict map[string]int = map[string]int{}
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        v := strings.Split(scanner.Text(), " ")
        for _, w := range v {
            if w != "" {
                if _, ok := dict[w]; ok {
                    dict[w]+=1
                } else {
                    dict[w] = 1
                }
            }
        }
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
    for key, val := range dict {
        fmt.Printf("%s\t%d\n", strings.TrimSpace(key), val)
    }
}
