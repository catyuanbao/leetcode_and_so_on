# run time 0.377

# 计算表达式的值
# 参数:
# numbers: 数字列表
# operators: 运算符列表 ('+' 或 '*')
def evaluate_expression(numbers, operators):
    # 初始化结果为第一个数字
    result = numbers[0]
    # 遍历操作符列表并依次计算结果
    for i in range(len(operators)):
        if operators[i] == '+':  # 如果操作符是加法
            result += numbers[i + 1]  # 将当前数字加到结果中
        elif operators[i] == '*':  # 如果操作符是乘法
            result *= numbers[i + 1]  # 将当前数字乘到结果中
    return result  # 返回最终计算结果

# 使用递归生成所有可能的运算符组合
# 参数:
# n: 需要的运算符数量 (数字数量 - 1)
# 返回:
# 包含所有运算符组合的列表
def generate_operators(n):
    # 基础情况: 如果不需要运算符，返回空列表
    if n == 0:
        return [[]]  # 空列表表示没有运算符
    # 递归生成 n-1 个运算符组合
    smaller = generate_operators(n - 1)
    result = []
    # 遍历较小规模的组合，并为每种组合追加'+'和'*'
    for ops in smaller:
        result.append(ops + ['+'])  # 添加加法组合
        result.append(ops + ['*'])  # 添加乘法组合
    return result  # 返回所有组合

# 解析输入数据并计算符合条件的总校准值
# 参数:
# input_data: 包含多行输入数据的列表，每行格式为 "目标值: 数字列表"
# 返回:
# 满足条件的目标值总和
def calculate_calibration(input_data):
    total_calibration = 0  # 初始化校准值总和

    # 遍历每一行输入数据
    for line in input_data:
        # 解析每一行数据: 获取目标值和数字列表
        target, values = line.split(': ')  # 按冒号拆分
        target = int(target)  # 将目标值转换为整数
        numbers = list(map(int, values.split()))  # 将数字字符串转换为整数列表

        # 生成所有可能的运算符组合
        operators_list = generate_operators(len(numbers) - 1)

        # 检查是否存在满足条件的组合
        for operators in operators_list:
            # 如果当前运算符组合计算的结果等于目标值
            if evaluate_expression(numbers, operators) == target:
                total_calibration += target  # 将目标值加入总和
                break  # 找到符合条件的组合后退出循环

    return total_calibration  # 返回总校准值

# 示例输入数据
# 测试样例: 每一行为目标值和数字列表
data = [
    "190: 10 19",
    "3267: 81 40 27",
    "83: 17 5",
    "156: 15 6",
    "7290: 6 8 6 15",
    "161011: 16 10 13",
    "192: 17 8 14",
    "21037: 9 7 18 13",
    "292: 11 6 16 20"
]

# 从文件读取输入数据
lines = []  # 存储文件内容
with open('data.in') as fd:  # 打开文件读取
    for line in fd:
        lines.append(line.strip())  # 去掉每行首尾空格并存入列表

# 使用文件中的数据进行计算
data = lines

# 计算结果并打印输出
result = calculate_calibration(data)
print("Total Calibration Result:", result)

