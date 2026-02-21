from typing import List

# Segment Tree
class NumArray:

    def __init__(self, nums: List[int]):
        self.nums = nums
        self.n = len(nums)
        self.segment_tree = []
        self.buildSegmentTree(0,0,self.n-1)

    def buildSegmentTree(self,i,l,r):
    	if l==r:
    		self.segment_tree[i]=self.nums[l]
    		return
    	mid = (l+(r-l))//2
    	self.buildSegmentTree(2*i+1,l,mid)
    	self.buildSegmentTree(2*i+2,mid+1,r)
    	self.segment_tree[i] = self.segment_tree[2*i+1]+self.segment_tree[2*i+2]

    def update(self, index: int, val: int) -> None:
        def update_helper(i,l,r):
        	if l==r==index:
        		self.segment_tree[i]=val
        		return
        	if l==r:
        		return
        	mid = (l+(r-l))//2
        	if index <= mid:
        		update_helper(2*i+1,l,mid)
        	else:
        		update_helper(2*i+2,mid+1,r)
        	self.segment_tree[i]=self.segment_tree[2*i+1]+self.segment_tree[2*i+2]

        update_helper(0,0,self.n-1)

    def sumRange(self, left: int, right: int) -> int:
    	def sumRange_helper(i,l,r):
    		"""
				First conditon if the range is fully exclusive
    		"""
    		if right < l or r < left:
    			return 0
    		elif left <= l <=r <= right:
    			return self.segment_tree[i]
    		else:
    			mid = (l+(r-l))//2
    			return sumRange_helper(2*i+1,l,mid)+sumRange_helper(2*i+2,mid+1,r)

    	return sumRange_helper(0,0,self.n-1)
    	

if __name__ == "__main__":
	nums:List[int] = [0,9,5,7,3]
	n = NumArray(nums)
	print(n.segment_tree)
	print(n.sumRange(0,2))
	n.update(1,2)
	print(n.segment_tree)
	print(n.sumRange(0,2))