package main

import (
        "bufio"
        "fmt"
        "os"
        "strconv"
)

func FindAns(nums []int64) int64 {
        var res int64
		numMap := make(map[int64][]int64)
        res = -1
		length := len(nums)
		for i := 0; i < length; i++ {
			for j := i; j < length; j++ {
				key := 2020 - nums[i]
				if _, ex := numMap[key]; ex {
					fmt.Println(nums[i], numMap[key])
					res = nums[i] * numMap[key][0] * numMap[key][1]
					return res
				}
				sum := nums[i] + nums[j]
				numMap[sum] = append(numMap[sum], nums[i])
				numMap[sum] = append(numMap[sum], nums[j])
			}
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

