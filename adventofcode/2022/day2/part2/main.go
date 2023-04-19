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
        result_map["A"] = 1
        result_map["B"] = 2
        result_map["C"] = 3

        mathch_result_map := make(map[string]int)
        mathch_result_map["X"] = 0
        mathch_result_map["Y"] = 3
        mathch_result_map["Z"] = 6

        score_map := make(map[string]map[string]string)
        var score_map_value_a map[string]string
        score_map_value_a = make(map[string]string)

        score_map_value_a["X"] = "C"
        score_map_value_a["Y"] = "A"
        score_map_value_a["Z"] = "B"

        score_map["A"] = score_map_value_a

        var score_map_value_b map[string]string
        score_map_value_b = make(map[string]string)

        score_map_value_b["X"] = "A"
        score_map_value_b["Y"] = "B"
        score_map_value_b["Z"] = "C"

        score_map["B"] = score_map_value_b

        var score_map_value_c map[string]string
        score_map_value_c = make(map[string]string)

        score_map_value_c["X"] = "B"
        score_map_value_c["Y"] = "C"
        score_map_value_c["Z"] = "A"

        score_map["C"] = score_map_value_c

        // fmt.Println(score_map)

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                line := scanner.Text()
                player_mark := strings.Split(line, " ")
                // fmt.Println(player_mark)
                player := player_mark[0]
                your := player_mark[1]
                total += result_map[score_map[player][your]] + mathch_result_map[your]
        }
        fmt.Println(total)

        if err := scanner.Err(); err != nil {
                fmt.Println(err)
        }
}

