import java.util.*;



class Solution{
	static class NumArray {

    private int[] nums;
    private int n;
    private int[] segmentTree;


    public NumArray(int[] nums) {
        this.nums = nums;
        this.n = nums.length;
        this.segmentTree = new int[4*n];
        this.buildSegmentTree(0,0,n-1);
    }

    private void buildSegmentTree(int i, int left,int right){
    	if(left == right){
    		this.segmentTree[i]=this.nums[left];
    		return ;
    	}

    	int mid = (right+left)/2;
    	System.out.println(left+" "+mid);
    	buildSegmentTree(2*i+1,left,mid);
    	System.out.println(left+" "+(mid+1)+" "+right);
    	buildSegmentTree(2*i+2,mid+1,right);
    	this.segmentTree[i] =this.segmentTree[2*i+1]+this.segmentTree[2*i+2];

    }
    
    public void update(int index, int val) {
     	   updateHelper(0,0,this.n-1,index,val);
    }

    private void updateHelper(int i,int l,int r,int index,int val){
    	if(l == r && r == index){
    		this.segmentTree[i]=val;
    		return;
    	}
    	if(l==r){
    		return ;
    	}

    	int mid = (r+l)/2;
    	if(index <= mid){
    		updateHelper(2*i+1,l,mid,index,val);
    	}else{
    		updateHelper(2*i+2,mid+1,r,index,val);
    	}
    	this.segmentTree[i]=this.segmentTree[2*i+1]+this.segmentTree[2*i+2];
    }
    
    public int sumRange(int left, int right) {
        return sumRangeHelepr(0,left,right,0,this.n-1);
    }

    private int sumRangeHelepr(int i,int left,int right,int l,int r){
    		System.out.println("left "+l+" right "+r);
    		if( right < l || r < left){
    			return 0;
    		}else if ( left<=l && r <= right){
    			return this.segmentTree[i];
    		}else{
    			int mid = (r+l)/2;
    			return sumRangeHelepr(2*i+1,left,right,l,mid)+sumRangeHelepr(2*i+2,left,right,mid+1,r);
    		}
    }
}
	public static void main(String[] args) {
		int [] nums = {9,-8};
		NumArray numArray = new NumArray(nums);
		// System.out.println(Arrays.toString(numArray.segmentTree));
		System.out.println(numArray.sumRange(1,1));
		numArray.update(0,3);
		// System.out.println(Arrays.toString(numArray.segmentTree));
		System.out.println(numArray.sumRange(1,1));
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * NumArray obj = new NumArray(nums);
 * obj.update(index,val);
 * int param_2 = obj.sumRange(left,right);
 */