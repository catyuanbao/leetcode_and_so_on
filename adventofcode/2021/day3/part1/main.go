package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FindAns(all_strings []string) int64 {
        length := len(all_strings[0])
	len_of_all_strings := len(all_strings)
	half_len := len_of_all_strings / 2
        most_bit := make([]int, length)

        for i := 0; i < length; i++ {
                for _, str := range all_strings {
                        if string(str[i]) == "1" {
                                most_bit[i] += 1
                        }
                }
        }

	most_bit_list := make([]string, 0)
	less_bit_list := make([]string, 0)

	for _, bit := range most_bit {
		if bit > half_len {
			most_bit_list = append(most_bit_list, "1")
			less_bit_list = append(less_bit_list, "0")
		} else {
			most_bit_list = append(most_bit_list, "0")
			less_bit_list = append(less_bit_list, "1")
		}
	}

	num_a, err := strconv.ParseInt(strings.Join(most_bit_list, ""), 2, 64)
	if err != nil {
		return -1
	}
	num_b, err := strconv.ParseInt(strings.Join(less_bit_list, ""), 2, 64)
	if err != nil {
		return -2
	}
        return num_a * num_b
}


func main() {
	file, err := os.Open("in.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	all_strings := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		all_strings = append(all_strings, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(FindAns(all_strings))
}
