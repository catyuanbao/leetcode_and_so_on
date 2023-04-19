package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func FindAns(all_lines []string) uint64 {
	var total uint64
	for _, line := range all_lines {
		half_length_of_str := len(line) / 2
		first_str, second_str := line[:half_length_of_str], line[half_length_of_str:]
		var common_char byte
		// TODO find better way to find common str
		for _, char := range first_str {
			for _, char2 := range second_str {
				if char == char2 {
					common_char = byte(char)
					break
				}
			}
		}
		str_common := string(common_char)
		if IsUpper(str_common) {
			num, _ := strconv.ParseUint(fmt.Sprintf("%d", common_char-'A'+27), 10, 64)
			total += num
		} else {
			num, _ := strconv.ParseUint(fmt.Sprintf("%d", common_char-'a'+1), 10, 64)
			total += num
		}
	}
	return total
}

func main() {
	file, err := os.Open("ex.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// fmt.Println(score_map)
	all_line := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		all_line = append(all_line, line)
	}

	fmt.Println(FindAns(all_line))

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
