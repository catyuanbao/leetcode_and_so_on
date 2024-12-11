package main

import (
        "bufio"
        "fmt"
        "os"
        "strconv"
)

func FindAns(nums []int64) int64 {
        var res int64
		numMap := make(map[int64]int64, len(nums))
        res = -1
        for _, value := range nums {
			toFind := 2020 - value
			if _, e := numMap[toFind]; e {
				fmt.Println("find:", value, toFind)
				return value * toFind
			}
			numMap[value] = 1
        }

        return res
}

func main() {
        file, err := os.Open("data.in")
        if err != nil {
                fmt.Println(err)
                return
        }
        defer file.Close()

        all_depth := make([]int64, 0)

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                line := scanner.Text()
                depth, err := strconv.ParseInt(line, 10, 64)
                if err == nil {
                        all_depth = append(all_depth, depth)
                }
        }
        // fmt.Println(all_depth)
        fmt.Println(FindAns(all_depth))

        if err := scanner.Err(); err != nil {
                fmt.Println(err)
        }
}

