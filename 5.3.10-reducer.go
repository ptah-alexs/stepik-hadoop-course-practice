package main

import (
    "bufio"
    "os"
    "fmt"
)

func main() {
    tempKey := ""
    dfval := 0.0
    dvert := ""
    dparts := ""
    dvrtx := ""
    dpr := 0.0
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        fmt.Sscanf(scanner.Text(),"%s\t%f\t{%s}", &dvert, &dfval, &dparts)
        if tempKey != "" && dvert != tempKey {
            dd := 0.02 + (0.9*dpr)
            fmt.Printf("%s\t%.3f\t{%s}\n", tempKey, dd, dvrtx)
            dpr = 0.0
            tempKey = dvert
            dvrtx = ""
            dvrtx = dparts[:len(dparts)-1]
            if dparts[:len(dparts)-1] != "" {
            } else {
                dpr += dfval
            }
        } else {
            
            if dparts[:len(dparts)-1] != "" {
                dvrtx = dparts[:len(dparts)-1]
            } else {
                dpr += dfval
            }
            tempKey = dvert
        }
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
    if tempKey != "" {
        dvrtx = dparts[:len(dparts)-1]
        dd := 0.02 + (0.9*dpr)
        fmt.Printf("%s\t%.3f\t{%s}\n", tempKey, dd, dvrtx)
    }
}
