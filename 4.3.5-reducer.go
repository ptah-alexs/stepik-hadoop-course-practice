package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
)

func main() {
    tempKey := ""
    sum := 0
    sVals := []string{}
    var dict map[string]int = map[string]int{}
    tnar := []string{}
    tvar := []string{}
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        sVals = strings.Split(scanner.Text(), "\t")
        tnar = append(tnar, sVals[0])
        tvar = append(tvar, sVals[1])
        if tempKey != "" && sVals[0] != tempKey {
            dict[tempKey] = sum
            tempKey = sVals[0]
            sum = 1
        } else {
            sum ++
            tempKey = sVals[0]
        }
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
    if tempKey != "" {
        dict[tempKey] = sum
    }
    for n, el := range tnar {
        parts := strings.Split(tvar[n],";")
        fmt.Printf("%s#%s\t%s\t%d\n", el, parts[0], parts[1], dict[el])
    }
}
