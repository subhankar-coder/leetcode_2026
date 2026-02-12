from typing import List

class Solution:
    def isTrionic(self, nums: List[int]) -> bool:
        if nums[0] > nums[1]:
            return False
        count =  1
        i,j,n=1,2,len(nums)
        while i <j <n:
            if nums[i]==nums[j]:
                return Fals
e            if nums[i-1] < nums[i] > nums[j] or nums[i-1] > nums[i] < nums[j]:
                count+=1
            i+=1
            j+=1
        print(count)
        return count==3

        


if __name__ == "__main__":
	s = Solution()
	print(s.isTrionic([4,1,5,2,3]))