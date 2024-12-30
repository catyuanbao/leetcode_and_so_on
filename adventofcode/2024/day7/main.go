package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// evaluateExpression 计算表达式的值
// 参数:
//   numbers - 数字列表
//   operators - 运算符列表 ('+' 或 '*')
// 返回:
//   计算结果
func evaluateExpression(numbers []int, operators []string) int {
	result := numbers[0] // 初始化结果为第一个数字
	for i := 0; i < len(operators); i++ {
		if operators[i] == "+" { // 如果操作符是加法
			result += numbers[i+1] // 执行加法运算
		} else if operators[i] == "*" { // 如果操作符是乘法
			result *= numbers[i+1] // 执行乘法运算
		}
	}
	return result // 返回计算结果
}

// generateOperators 递归生成所有可能的运算符组合
// 参数:
//   n - 需要的运算符数量 (数字数量 - 1)
// 返回:
//   包含所有运算符组合的二维字符串切片
func generateOperators(n int) [][]string {
	if n == 0 { // 基础情况
		return [][]string{{}} // 返回空组合
	}

	// 递归生成 n-1 的组合
	smaller := generateOperators(n - 1)
	var result [][]string
	for _, ops := range smaller {
		// 在每个组合后追加 '+' 和 '*'
		result = append(result, append([]string{}, append(ops, "+")...))
		result = append(result, append([]string{}, append(ops, "*")...))
	}
	return result
}

// calculateCalibration 解析输入数据并计算符合条件的总校准值
// 参数:
//   inputData - 包含多行输入数据的字符串切片，每行格式为 "目标值: 数字列表"
// 返回:
//   符合条件的目标值总和
func calculateCalibration(inputData []string) int {
	totalCalibration := 0 // 初始化总和

	for _, line := range inputData {
		// 按 ": " 拆分目标值和数字列表
		parts := strings.Split(line, ": ")
		target, _ := strconv.Atoi(parts[0])       // 转换目标值为整数
		numStrs := strings.Fields(parts[1])      // 拆分数字列表
		numbers := make([]int, len(numStrs))     // 创建整数切片
		for i, num := range numStrs {
			numbers[i], _ = strconv.Atoi(num) // 转换数字为整数
		}

		// 生成所有运算符组合
		operatorsList := generateOperators(len(numbers) - 1)

		// 检查是否存在满足条件的组合
		for _, operators := range operatorsList {
			if evaluateExpression(numbers, operators) == target { // 校验结果是否匹配目标值
				totalCalibration += target
				break // 找到匹配的组合后退出
			}
		}
	}

	return totalCalibration // 返回总和
}

// 主程序入口
// 运行时间: 0.232
func main() {
	// 从文件读取输入数据
	file, err := os.Open("data.in")
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // 按行读取数据
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件错误:", err)
		return
	}

	// 计算校准值并输出结果
	result := calculateCalibration(data)
	fmt.Println("Total Calibration Result:", result)
}

