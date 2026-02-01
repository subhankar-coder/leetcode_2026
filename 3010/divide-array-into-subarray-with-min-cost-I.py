from typing import List

class Solution:
    def minimumCost(self, nums: List[int]) -> int:
        part = nums[1:]
        part.sort()
        return nums[0]+part[0]+part[1]
        
        

if __name__ == "__main__":
    s = Solution()
    print(s.minimumCost([1,2,3,12]))