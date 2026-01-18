package main

import(
	"fmt"
)

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
    var maxArea int64 = 0
    for i:= range bottomLeft{
    	for j:=i+1;j<len(bottomLeft);j++{
    		left,right,bottom,top := max(bottomLeft[i][0],bottomLeft[j][0]),
    		min(topRight[i][0],topRight[j][0]),max(bottomLeft[i][1],bottomLeft[j][1]),
    		min(topRight[i][1],topRight[j][1])
    		if left < right && bottom < top{
    			x,y := right-left , top-bottom
    			area := int64(min(x,y)*min(x,y))
    			maxArea = max(maxArea,area)
    		}
    	}
    }
    // fmt.Println(maxArea)
    return maxArea
}

func main(){
	bottomLeft,topRight := [][]int{{2,2},{2,3}},[][]int{{4,3},{3,4}}
	fmt.Println(largestSquareArea(bottomLeft,topRight))
}