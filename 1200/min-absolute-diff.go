package main

import(
	"fmt"
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
    mp,min_diff,n := make(map[int][][]int),math.MaxInt,len(arr)
    sort.Ints(arr)
    for i,j:=0,1;i<n-1 && j<n;i,j=i+1,j+1{
    	// for j:=i+1;j<len(arr);j++{
    		diff := int(math.Abs(float64(arr[j]-arr[i])))
    		mp[diff]=append(mp[diff],[]int{arr[i],arr[j]})
    		min_diff = min(min_diff,diff)
    	// }
    }
    // // fmt.Println(min_diff)
    //  sort.Slice(mp[min_diff], func(i,j int) bool{
    // 	if mp[min_diff][i][0]!=mp[min_diff][j][0]{
    // 		return mp[min_diff][i][0] < mp[min_diff][j][0]
    // 	}
    // 	return mp[min_diff][i][1] < mp[min_diff][j][1]
    // })

    return mp[min_diff]
}

func main(){
	arr := []int{4,2,1,3}
	fmt.Println(minimumAbsDifference(arr))
}