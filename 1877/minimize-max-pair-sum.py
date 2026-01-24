from typing import List


class Solution:
    def minPairSum(self, nums: List[int]) -> int:
    	nums.sort()
    	start,end,mx = 0,len(nums)-1,float("-inf")
    	while start<end:
    		mx = max(mx,nums[start]+nums[end])
    		start+=1
    		end-=1
    	return mx

if __name__ == "__main__":
	s = Solution()
	nums:List = [3,5,4,2,4,6]
	print(s.minPairSum(nums))