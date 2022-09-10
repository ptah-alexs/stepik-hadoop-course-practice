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
        sVals := strings.Split(scanner.Text(), " ")
        for i:=0; i<len(sVals); i++ {
            for j:=0; j<len(sVals); j++ {
                if sVals[i] != sVals[j] {
                    fmt.Printf("%s,%s\t1\n",sVals[i],sVals[j])
                }
            }
        }
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
}

