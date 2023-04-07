package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func find_max(sum_list []int64) int64 {
    var ret int64 
    ret = 0
    for _, ele := range sum_list {
        if ele >= ret {
            ret = ele
        }
    }
    return ret
}


func main() {
    file, err := os.Open("in.in")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var sum_one int64
    // one Elve carry total Calories
    sum_one = 0

    sum_list := make([]int64, 0)
    for scanner.Scan() {
        line := scanner.Text()

        one, err := strconv.ParseInt(line, 10, 64)
        // parse new line not a number, append current sum_one and reset sum_one to 0
        if err != nil {
            sum_list = append(sum_list, sum_one)
            sum_one = 0
        }
        sum_one += one
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }

    fmt.Println(find_max(sum_list))
}

