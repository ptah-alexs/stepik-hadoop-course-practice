package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "regexp"
)

func main() {
    tempKey := ""
    file1 := []string{}
    re := regexp.MustCompile(`\w+`)
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        sVals := strings.SplitN(scanner.Text(), ":", 2)
        file1 = re.FindAllString(sVals[1], -1)
        if tempKey != "" && sVals[0] != tempKey {
            for _, el := range file1 {
                fmt.Printf("%s#%s\t1\n", el, sVals[0])
            }
            tempKey = sVals[0]
            file1 = []string{}
        } else {
            for _, el := range file1 {
                fmt.Printf("%s#%s\t1\n", el, sVals[0])
            }
            tempKey = sVals[0]
        }
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
}
