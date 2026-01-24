import java.util.*;

class Solution{
	static class Node{
		Node left,right;
		long value;
		int index;
		public Node(long value,int index){
			this.value=value;
			this.index = index;
		}
	}
	record PQItems(Node first,Node second,long cost){}
	public static int minimumPairRemoval(int[] nums) {
        int n = nums.length;
        Node head, current =new Node(nums[0],0);
        boolean [] merged = new boolean[n];
        int differenceCount=0;
        int count =0;
        PriorityQueue<PQItems> pq = new PriorityQueue<>(
        	Comparator.comparingLong(PQItems::cost)
        	.thenComparingInt(p->p.first().index)
        );
        for(int i=1;i<n;i++){
        	Node newNode = new Node(nums[i],i);
        	current.right = newNode;
        	newNode.left = current;
        	pq.offer(new PQItems(current,newNode,current.value+newNode.value));
        	if(current.value > newNode.value) differenceCount++;
        	current = newNode;
        }
        // System.out.println(differenceCount);
        while(differenceCount >0){
        	PQItems item = pq.poll();
        	Node first = item.first;
        	Node second = item.second;
        	long cost = item.cost;
        	if(merged[first.index] || merged[second.index] || first.value+second.value != cost) continue;
        	count++;
        	if (first.value > second.value) differenceCount--;
        	Node prevFirst = first.left;
        	Node nextSecond = second.right;

        	first.right = nextSecond;
        	if(nextSecond!=null) nextSecond.left = first;
        	if (prevFirst!=null){
        		if(prevFirst.value > first.value && prevFirst.value <= cost) differenceCount--;
        		else if(prevFirst.value <= first.value && prevFirst.value > cost) differenceCount++;
        		pq.offer(new PQItems(prevFirst,first,prevFirst.value+cost));
        	}
        	if(nextSecond!=null){
        		if(second.value > nextSecond.value && cost <= nextSecond.value) differenceCount--;
        		else if(second.value <= nextSecond.value && cost > nextSecond.value) differenceCount++;
        		pq.offer(new PQItems(first,nextSecond,cost+nextSecond.value));
        	}
        	first.value = cost;
        	merged[second.index]=true;
        }

        return count;
    }

	public static void main(String[] args) {
		int [] nums = {2,2,-1,3,-2,2,1,1,1,0,-1};
		System.out.println(minimumPairRemoval(nums));
	}
}