package main

import (
	"fmt"
)

func largestMagicSquare(grid [][]int) int {
	m,n := len(grid),len(grid[0])
	maxSize,maxResult := min(m,n),1
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

	// fmt.Println(rowSum)
	// fmt.Println(colSum)
	for size:=2;size<=maxSize;size++{
		for row:=0;row<m-size+1;row++{
			for col:=0;col<n-size+1;col++{
				// fmt.Println(row,col,size)
				rowAgg,colAgg,flag,currRowSum,currColSum,diag1,diag2,diagx1,diagx2 := 0,0,false,0,0,0,0,col,col+size-1
				for start:= row;start<row+size;start++{
					if col == 0{
						currRowSum = rowSum[start][col+size-1]
					}else{
						currRowSum = rowSum[start][col+size-1]-rowSum[start][col-1]
					}
					if rowAgg !=0 && rowAgg!=currRowSum{
						// fmt.Println(row,col,currRowSum)
						flag = true
						break
					}else{
						rowAgg = currRowSum
					}
				}
				if flag{
					continue
				}

				for start := col;start<col+size;start++{
					if row ==0{
						currColSum = colSum[row+size-1][start]
					}else{
						currColSum = colSum[row+size-1][start]-colSum[row-1][start]
					}
					if colAgg !=0 && colAgg != currColSum{
						flag = true
						break
					}else{
						colAgg = currColSum
					}
				}

				if flag || rowAgg != colAgg{
					continue
				}
				for start := row ; start<row+size; start++{
					diag1 += grid[start][diagx1]
					diag2 += grid[start][diagx2]
					diagx1+=1
					diagx2-=1
				}
				if diag1 == diag2 && diag1 == rowAgg{
					maxResult = max(maxResult,size)
				}
			}
		}
	}
	return maxResult
}

func main(){
	grid := [][]int{{5,1,3,1},{9,3,3,1},{1,3,3,8}};
	fmt.Println(largestMagicSquare(grid))
}
