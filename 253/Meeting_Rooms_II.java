import java.util.*;

class Solution {
    public int minMeetingRooms(int[] start, int[] end) {
        TreeMap<Integer,Integer> mp = new TreeMap<>();
        int n = start.length,max_result=0,total=0;
		for(int i=0;i<n;i++){
		    int s = mp.getOrDefault(start[i],0)+1;
		    int e = mp.getOrDefault(end[i],0)-1;
			mp.put(start[i],s);
			mp.put(end[i],e);
		}
		
		for(Integer e: new ArrayList<Integer>(mp.values())){
		    total+=e;
		    max_result = Math.max(max_result,total);
		}
		return max_result;
        
    }
}
