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
    tnar := []string{}
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        sVals = strings.Split(scanner.Text(), "\t")
        if tempKey != "" && sVals[0] != tempKey {
            for _, el := range tnar {
                fmt.Printf("%s%d\n",el, sum)
            }
            tempKey = sVals[0]
            sum = 1
            tnar = []string{}
            parts := strings.Split(sVals[1],";")
            tnar = append(tnar, sVals[0] + "#" + parts[0] + "\t" + parts[1] +"\t")
        } else {
            parts := strings.Split(sVals[1],";")
            tnar = append(tnar, sVals[0] + "#" + parts[0] + "\t" + parts[1] +"\t")
            sum ++
            tempKey = sVals[0]
        }
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
    if tempKey != "" {
        for _, el := range tnar {
            fmt.Printf("%s%d\n",el, sum)
        }
    }
}
