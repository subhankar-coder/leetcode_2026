package main

import(
	"fmt"
)

func maxSideLength(mat [][]int, threshold int) int {
    m,n,start,result := len(mat),len(mat[0]),0,0
    end := min(m,n)
    rowSum := make([][]int,m)
    for i:= range m{
    	rowSum[i]=make([]int,n)
    	for j:= range n{
    		if j==0{
    			rowSum[i][j] = mat[i][0]
    		}else{
    			rowSum[i][j] = rowSum[i][j-1]+mat[i][j]
    		}
    	}
    }
    isPossible := func(size int )bool{
    	for row := 0;row <= m-size;row++ {
    		for col := 0;col <= n-size;col++ {
    			total_sum := 0
    			for r:=row;r<row+size;r++{
    				sum :=0
    				if col ==0{
    					sum = rowSum[r][size-1]
    				}else{
    					sum = rowSum[r][size+col-1]-rowSum[r][col-1]
    				}
    				total_sum+=sum
    			}
    			if total_sum > threshold{
    				continue
    			}else{
    				return true
    			}
    		}
    	}
    	return false
    }
    for start <= end{
    	mid := (start+end)/2
    	if isPossible(mid){
    		result = mid
    		start = mid+1
    	}else{
    		end = mid -1
    	}
    }

    return result
}

func main(){
	mat,threshold := [][]int{{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2}},1
	fmt.Println(maxSideLength(mat,threshold))
}