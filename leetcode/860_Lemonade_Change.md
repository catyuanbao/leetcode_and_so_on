```go

func lemonadeChange(bills []int) bool {
	money_have := make([]int, 2)
	for _, bill := range bills {
		// input 5
		if bill == 5 {
			money_have[0] += 1
		} else if bill == 10 {
			money_have[1] += 1
			if money_have[0] >= 1 {
				money_have[0] -= 1
			} else {
				return false
			}
		} else {
			if money_have[0] >= 1 && money_have[1] >= 1 {
				money_have[0] -= 1
				money_have[1] -= 1
			} else if money_have[0] >= 3 {
				money_have[0] -= 3
			} else {
				return false
			}
		}
	}
	return true
}

```
