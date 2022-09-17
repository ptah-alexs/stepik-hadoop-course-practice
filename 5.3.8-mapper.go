package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
)

func main() {
    dval := 0.0
    dvert := ""
    dparts := ""
    nparts := 0
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
        fmt.Sscanf(scanner.Text(),"%s\t%f\t{%s}", &dvert, &dval, &dparts)
        tstr := dparts[:len(dparts)-1]
        sparts :=  strings.Split(tstr, ",")
        nparts = len(sparts)
        for _, el := range sparts {
            fmt.Printf("%s\t%.3f\t{}\n", el, dval/float64(nparts))
        }
        
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
}
