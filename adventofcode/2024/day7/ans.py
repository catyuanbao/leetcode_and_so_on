a = '''
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
'''

# run time  1.426 total

lines = []
with open('data.in') as fd:
    for line in fd:
        lines.append(line.strip())

import itertools
sum_all = 0
for line in lines:
    nums = line.replace(':', '').split(' ')
    target = int(nums[0])
    factor = nums[1:]
    length = len(factor[1:])
    operators = list(itertools.product('+*', repeat=length))
    find_line = False

    for i in operators:
        strop = ''.join(i)
        length_nums = len(factor[1:])
        index = 0
        temp_res = None
        while index < length:
            if temp_res is None:
                if strop[index] == '+':
                    temp_res = int(factor[index]) + int(factor[index+1])
                elif strop[index] == '*':
                    temp_res = int(factor[index]) * int(factor[index+1])
            else:
                if strop[index] == '+':
                    temp_res = temp_res + int(factor[index+1])
                if strop[index] == '*':
                    temp_res = temp_res * int(factor[index+1])
            if temp_res == target:
                sum_all += target
                find_line = True
                break
            index += 1

        if find_line:
            break

print("AAA", sum_all)

