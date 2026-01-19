package main

import "fmt"

func largestMagicSquare(grid [][]int) int {
	m,n := len(grid),len(grid[0])
	maxSize := min(m,n)
	rowSum := make([][]int,m)
	colSum := make([][]int,m)
	for i:= range m{
		rowSum[i] = make([]int, n)
		colSum[i]= make([]int, n)
	}
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if j==0{
				rowSum[i][j] = grid[i][j]
			}else{
				rowSum[i][j] = rowSum[i][j-1]+grid[i][j]
			}
		}
	}
	
	for i:=0;i<n;i++{
		for j:=0;j<m;j++{
			if j==0{
				colSum[j][i] = grid[j][i]
			}else{
				colSum[j][i] = colSum[j-1][i]+grid[j][i]
			}
		}
	}

	fmt.Println(rowSum)
	fmt.Println(colSum)
	for size:=2;size<=maxSize;size++{
		for row:=0;row<=m-size;row++{
			for col:=0;col<=n-size;col++{
				
			}
	}
	return 0
}

func main(){
	grid := [][]int{{7,1,4,5,6},{2,5,1,6,4},{1,5,4,3,2},{1,2,7,3,4}};
	fmt.Println(largestMagicSquare(grid))
}
