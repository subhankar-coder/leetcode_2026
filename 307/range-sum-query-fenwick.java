import java.util.*;

public class Solution{
	public static void main(String[] args) {
		int [] nums = {1, 3, 5};
		NumArray numArray = new NumArray(nums);
		System.out.println(Arrays.toString(numArray.fenW));
		System.out.println(numArray.sumRange(0,2));
		// numArray.update(1,2);
		System.out.println(numArray.sumRange(0,2));
		System.out.println(numArray.lowerBoundOnPrefixSum(4));
	}
	static class NumArray {
		private int[] fenW;
		private int N;
		private int[] nums;
	    public NumArray(int[] nums) {
	        this.nums = nums;
	        this.N = nums.length;
	        this.fenW = new int[N+1];
	        buildFenWickTree();
	    }

	    private void buildFenWickTree(){
	    	for(int i=0;i<this.N;i++){
	    		// this.fenW[i+1]+=this.nums[i];
	    		int next = i+1;
	    		while(next <=this.N){
	    			this.fenW[next]+=this.nums[i];
	    			next = next + (next & -next);
	    		}
	    	}
	    }
	    
	    public void update(int index, int val) {
	        int diff = val-this.nums[index];
	        this.nums[index]=val;
	        int i = index+1;
	        while (i <=N){
	        	this.fenW[i]+=diff;
	        	i = i+(i&-i);
	        }
	    }
	    public int sum(int i){
	    	int s =0;
	    	while(i>0){
	    		s+=this.fenW[i];
	    		i = i - (i&-i);
	    		// System.out.println("i "+i);
	    	}
	    	return s;	
	    }
	    public int sumRange(int left, int right) {
	        return sum(right+1)-sum(left);

	    }

	    public int lowerBoundOnPrefixSum(int target){
	    	int curr=1,prevSum=0;
	    	System.out.println(log2(this.N));
	    	for(int i=log2(this.N);i>=0;i--){
	    		// System.out.println(i+" "+(curr+(1<<i))+" "+this.fenW[curr+(1<<i)]);
	    		if(sum(curr+(1<<i))+prevSum < target){
	    			curr = curr+(1<<i);
	    			prevSum+=sum(curr+(1<<i));

	    		}
	    	}
	    	return curr+1;
	    }

	    private int log2(int val){
	    	return 31-Integer.numberOfLeadingZeros(val);
	    }
	}

}