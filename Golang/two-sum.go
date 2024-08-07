class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        tmp_dict = dict()
        for i, num in enumerate(nums):
            if num in tmp_dict:
                return [tmp_dict[num], i]
            else:
                tmp_dict[target - num] = i
