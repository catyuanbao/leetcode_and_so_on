```Python

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        occured = {}
        for index,value in enumerate(nums):
            to_find = target - value
            // if find answser return it else update dict
            if to_find in occured:
                return [occured[to_find], index]
            else:
                occured[value] = index

        return [0, 0]

```
