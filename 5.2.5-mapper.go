package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "strconv"
)

func main() {
    distDict := map[string]string{}
    dsval := ""
    dvert := ""
    dparts := ""
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
        fmt.Sscanf(scanner.Text(),"%s\t%s\t{%s}", &dvert, &dsval, &dparts)
        distDict[dvert] = dsval
        if dparts[:len(dparts)-1] != "" {
            sparts :=  strings.Split(dparts[:len(dparts)-1], ",")
            for _, val := range sparts {
                st := "INF"
                if distDict[dvert] != "INF" {
                    ff := 0
                    fmt.Sscanf(distDict[dvert], "%d", &ff)
                    st = strconv.Itoa(ff+1)
                }
                fmt.Printf("%s\t%s\t{}\n", val, st)
            }
        }
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
}
