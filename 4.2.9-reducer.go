package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
)

func main() {
    tempKey := ""
    file1 := []string{}
    file2 := []string{}
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        sVals := strings.Split(scanner.Text(), "\t")
        if tempKey != "" && sVals[0] != tempKey {
            if len(file1) > 0 && len(file2) > 0 {
                for _, el1 := range file1 {
                    for _, el2 := range file2 {
                        fmt.Printf("%s\t%s\t%s\n", tempKey, el1, el2)
                    }
                }
            }
            tempKey = sVals[0]
            file1 = []string{}
            file2 = []string{}
            if strings.HasPrefix(sVals[1], "query:") {
                stn1 := strings.TrimPrefix(sVals[1], "query:")
                file1 = append(file1, stn1)
            }
            if  strings.HasPrefix(sVals[1], "url:") {
                stn2 := strings.TrimPrefix(sVals[1], "url:")
                file2 = append(file2, stn2)
            }
        } else {
            if strings.HasPrefix(sVals[1], "query:") {
                stn1 := strings.TrimPrefix(sVals[1], "query:")
                file1 = append(file1, stn1)
            }
            if  strings.HasPrefix(sVals[1], "url:") {
                stn2 := strings.TrimPrefix(sVals[1], "url:")
                file2 = append(file2, stn2)
            }
            tempKey = sVals[0]
        }
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
    if tempKey != "" {
        if len(file1) > 0 && len(file2) > 0 {
            for _, el1 := range file1 {
                for _, el2 := range file2 {
                    fmt.Printf("%s\t%s\t%s\n", tempKey, el1, el2)
                }
            }
        }
    }
}
