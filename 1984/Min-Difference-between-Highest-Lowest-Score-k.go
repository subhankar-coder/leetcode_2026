package main

import(
	"fmt"
	"sort"
)

func minimumDifference(nums []int, k int) int {
    sort.Ints(nums)
    n,min_diff := len(nums),10_000_000_000
    for i,j := 0,k-1;i<=n-k && j<n;i,j = i+1,j+1{
    	min_diff = min(min_diff,nums[j]-nums[i])
    }
    return min_diff
}

func main(){
	nums,k := []int{9,4,1,7},2
	fmt.Println(minimumDifference(nums,k))
}