import java.util.*;

class Solution{
	static boolean canAttend(int[][] arr) {
		TreeMap<Integer,Integer> mp = new TreeMap<>();
		for (int [] ele : arr){
			int start = mp.getOrDefault(ele[0],0)+1;
			int end = mp.getOrDefault(ele[1],0)-1;
			mp.put(ele[0],start);
			mp.put(ele[1],end);
		}
		int total=0;
		for(Map.Entry<Integer,Integer> e: mp.entrySet()){
			total+=e.getValue();
			if(total >1){
				return false;
			}
		}
     	return true;
    }
	public static void main(String[] args) {
		int [][] arr = {{1,4},{10,15},{7,10}};
		System.out.println(canAttend(arr));
	}
}