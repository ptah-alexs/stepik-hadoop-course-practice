package main

import (
    "bufio"
    "os"
    "fmt"
    "strconv"
)

func main() {
    tempKey := ""
    dsval := ""
    dvert := ""
    dparts := ""
    ddist := 900000000000000
    dtmp := 900000000000000
    dvrtx := ""
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        fmt.Sscanf(scanner.Text(),"%s\t%s\t{%s}", &dvert, &dsval, &dparts)
        dtmp = 900000000000000
        if tempKey != "" && dvert != tempKey {
            st := "INF"
            if ddist != 900000000000000 {
                st = strconv.Itoa(ddist)
            }
            fmt.Printf("%s\t%s\t{%s}\n", tempKey, st, dvrtx)
            tempKey = dvert
            ddist = 900000000000000
            dvrtx = ""
            if dsval != "INF" {
                fmt.Sscanf(dsval, "%d", &dtmp)
            }
            if dtmp < ddist {
                ddist = dtmp
            }
        } else {
            if dparts[:len(dparts)-1] != "" {
            dvrtx = dparts[:len(dparts)-1]
            }
            if dsval != "INF" {
                fmt.Sscanf(dsval, "%d", &dtmp)
            }
            if dtmp < ddist {
                ddist = dtmp
            }
            tempKey = dvert
        }
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
    if tempKey != "" {
            st := "INF"
            if ddist != 900000000000000 {
                st = strconv.Itoa(ddist)
            }
        fmt.Printf("%s\t%s\t{%s}\n", tempKey, st, dvrtx)
    }
}
