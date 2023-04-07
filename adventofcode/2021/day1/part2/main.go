package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func FindAns(depths []int64) int64 {
	var count int64
	count = 0
	for index, _ := range depths {
		// avoid index error 
		if (index == 0 || index > len(depths) - 3) {
			continue
		}
		cur_depth_sum := depths[index - 1] + depths[index] + depths[index + 1]
		next_depth_sum := depths[index] + depths[index + 1] + depths[index + 2]
		if (next_depth_sum > cur_depth_sum) {
			count += 1
		}
	}
	return count
}

func main() {
	file, err := os.Open("in.in")
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

