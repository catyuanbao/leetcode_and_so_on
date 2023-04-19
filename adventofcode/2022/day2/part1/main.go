package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	file, err := os.Open("ex.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	total := 0

	result_map := make(map[string]int)
	result_map["X"] = 1
	result_map["Y"] = 2
	result_map["Z"] = 3

	score_map := make(map[string]map[string]int)
	var score_map_value_a map[string]int
	score_map_value_a = make(map[string]int)

	score_map_value_a["X"] = 3
	score_map_value_a["Y"] = 6
	score_map_value_a["Z"] = 0

	score_map["A"] = score_map_value_a

	var score_map_value_b map[string]int
	score_map_value_b = make(map[string]int)

	score_map_value_b["X"] = 0
	score_map_value_b["Y"] = 3
	score_map_value_b["Z"] = 6

	score_map["B"] = score_map_value_b

	var score_map_value_c map[string]int
	score_map_value_c = make(map[string]int)

	score_map_value_c["X"] = 6
	score_map_value_c["Y"] = 0
	score_map_value_c["Z"] = 3

	score_map["C"] = score_map_value_c

	// fmt.Println(score_map)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		player_mark := strings.Split(line, " ")
		// fmt.Println(player_mark)
		player := player_mark[0]
		your := player_mark[1]
		total += score_map[player][your] + result_map[your]
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

