package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Action struct {
	Direction string
	Value     int64
}

func FindAns(all_actions []Action) int64 {
	var up, forword int64
	up = 0
	forword = 0
	for _, action := range all_actions {
		if action.Direction == "up" {
			up += action.Value
		} else if action.Direction == "down" {
			up -= action.Value
		} else {
			forword += action.Value
		}
	}
	if up < 0 {
		up = -up
	}
	return up * forword
}

func main() {
	file, err := os.Open("in.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	all_actions := make([]Action, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		direction_and_value := strings.Split(line, " ")
		value, err := strconv.ParseInt(direction_and_value[1], 10, 64)
		if err == nil {
			action := Action{direction_and_value[0], value}
			all_actions = append(all_actions, action)
		}
	}
	// fmt.Println(all_actions)
	fmt.Println(FindAns(all_actions))

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
