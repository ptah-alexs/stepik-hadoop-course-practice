package main

import (
    "bufio"
    "os"
    "fmt"
)

var vertexEdges = [][]int{}
var vertexEdgeWeights = [][]int{}
var vertexVisited = []bool{}
var vertexDistance = []int{}

func findMin(vd []int, vv []bool) (int, int) {
    min := vd[0]
    minindex := 0
    for n, el := range vv {
        if !el {
            min = vd[n]
            minindex = n
            break
        }
    }
    for n, el := range vd {
        if el < min && !vv[n] {
           min = el
            minindex = n
        }
    }
    return minindex, min
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    count := 0
    vertex := 0
    edges := 0
    startEdge := 0
    endEdge := 0
    for scanner.Scan() {
        tmpNum := 0
        tmpToNum := 0
        tmpWeight := 0
        line := scanner.Text()
        if count == 0 {
            fmt.Sscanf(line, "%d\t%d", &vertex, &edges)
            vertexEdges = make([][]int, vertex)
            vertexEdgeWeights = make([][]int, vertex)
        } else if count == edges+1 {
            fmt.Sscanf(line, "%d\t%d", &startEdge, &endEdge)
        } else {
            fmt.Sscanf(line, "%d\t%d\t%d", &tmpNum, &tmpToNum, &tmpWeight)
            vertexEdges[tmpNum-1] = append(vertexEdges[tmpNum-1], tmpToNum)
            vertexEdgeWeights[tmpNum-1] = append(vertexEdgeWeights[tmpNum-1], tmpWeight)
        }
        count++
    }
    if scanner.Err() != nil {
        fmt.Printf("error: %s\n", scanner.Err())
    }
    vertexVisited := make([]bool, vertex)
    vertexDistance := make([]int, vertex)
    for n, _ := range vertexDistance {
        vertexDistance[n] = 900000000000000
        vertexVisited[n] = false
    }
    vertexDistance[startEdge-1] = 0
    for i:=0; i < vertex; i++ {
        v, _ := findMin(vertexDistance, vertexVisited)
        for n, edgs := range vertexEdges[v] {
            if vertexDistance[edgs-1] > vertexDistance[v] + vertexEdgeWeights[v][n] {
                vertexDistance[edgs-1] = vertexDistance[v] + vertexEdgeWeights[v][n]
            }
        }
        vertexVisited[v] = true
    }
     if vertexDistance[endEdge-1] != 900000000000000 {
        fmt.Println(vertexDistance[endEdge-1])
    } else {
        fmt.Println(-1)
    }
}
