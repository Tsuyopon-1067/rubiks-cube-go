package main

import (
    "fmt"
    "strings"
)

func main() {
    var qube [][]int
    for i := 1; i <= 6; i++ {
        var tmp []int
        for j := 1; j <= 9; j++ {
            tmp = append(tmp, i*10 + j)
        }
        qube = append(qube, tmp)
    }

    printQube(qube)
}

func getCube(x int) string {
    var num string = fmt.Sprintf("%d", x%10)
    var col int = x / 10
    return colorStr(num, col)
}
func colorStr(s string, c int) string {
    colors := map[int]string{}
    colors[0] = "\033[49m"
    colors[1] = "\033[48;5;07m"
    colors[2] = "\033[48;5;01m"
    colors[3] = "\033[48;5;04m"
    colors[4] = "\033[48;5;208m"
    colors[5] = "\033[48;5;02m"
    colors[6] = "\033[48;5;03m"

    var arr[] string
    arr = append(arr, colors[c])
    arr = append(arr, s)
    arr = append(arr, colors[0])
    return strings.Join(arr, "")
}

func printQube(qube [][]int) {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            fmt.Print(" ")
        }
        for j := 0; j < 3; j++ {
            q := qube[0][i*3+j]
            fmt.Print(getCube(q))
        }
        fmt.Println()
    }

    for i := 0; i < 3; i++ {
        for col := 1; col <= 4; col++ {
            for j := 0; j < 3; j++ {
                q := qube[col][i*3+j]
                fmt.Print(getCube(q))
            }
        }
        fmt.Println()
    }

    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            fmt.Print(" ")
        }
        for j := 0; j < 3; j++ {
            q := qube[5][i*3+j]
            fmt.Print(getCube(q))
        }
        fmt.Println()
    }
}
