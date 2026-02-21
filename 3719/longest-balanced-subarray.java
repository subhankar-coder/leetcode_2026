import java.util.*;


public class Solution{
	public static int longestBalanced(int[] nums) {
        int start=0;
        int end = nums.length;
        int result = 0;
        for(int i=start;i<end;i++){
        	Map<Integer,Integer> mp = new HashMap<>();
        		int even=0;
        		int odd =0;

        	for(int j = i;j<end;j++){
        		        		
        			if(nums[j]%2 ==0 && !mp.containsKey(nums[j])){
        				even+=1;
        			}else if(!mp.containsKey(nums[j])){
        				odd+=1;
        			}
        			mp.put(nums[j],mp.getOrDefault(nums[j],0)+1);
        			
        		
        		if(odd == even) result = Math.max(result,(j-i+1));
        	}
        }

        return result;
    }

	public static void main(String[] args) {
		int [] nums = {2,5,4,3};
		System.out.println(longestBalanced(nums));
	}
}