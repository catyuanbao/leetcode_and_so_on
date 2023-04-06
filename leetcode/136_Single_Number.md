based on:

```python
>>> 0 ^ 0
0
>>> 100 ^ 0
100
>>> 100 ^ 100
0
>>> 
```

ans:

```go
func singleNumber(nums []int) int {
    ret := 0
    for _, num := range nums {
        ret ^= num
    }
    return ret
}
```

```c
int singleNumber(int* nums, int numsSize){
    int res = 0;
    for (int i = 0; i < numsSize; i++) {
        res ^= nums[i];
    }
    return res;
}
```

