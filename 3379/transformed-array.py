from typing import List

class Solution:
    def constructTransformedArray(self, nums: List[int]) -> List[int]:
        n = len(nums)
        result = [0]*n

       	for i in range(n):
       		if nums[i]==0:
       			result[i]=nums[i]
       		elif nums[i]>0:
       			result[i]=nums[(i+nums[i]) % n]
       		else:
       			ele = i-abs(nums[i])
       			result[i]=nums[((i+nums[i])%n+n)%n]

       	return result

if __name__ == "__main__":
	s = Solution()
	print(s.constructTransformedArray([-10,-10]))