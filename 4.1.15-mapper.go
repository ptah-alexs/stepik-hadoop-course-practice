package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "strconv"
)

func main() {
    var dict map[string]int = map[string]int{}
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        sVals := strings.Split(scanner.Text(), " ")
        for i:=0; i<len(sVals); i++ {
            dict = map[string]int{}
             ss := []string{}
            for j:=0; j<len(sVals); j++ {
                if sVals[i] != sVals[j] {
                    dict[sVals[j]]++
                }
            }
            for key, val := range dict {
                st := key+":"+strconv.Itoa(val)
                ss = append(ss,st)
            }
            fmt.Printf("%s\t%s\n",sVals[i], strings.Join(ss,","))
        }
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
}

