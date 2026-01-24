package main


import(
	"fmt"
	"math"
	// "sort"
)


func separateSquares(squares [][]int) float64 {
    total_area,max_height:=0.0,math.MinInt

    for i := range squares{
    	y,l := squares[i][1],squares[i][2]
    	total_area += float64(l*l)
    	max_height = max(max_height,y+l)
    }
    // sort.Slice(squares,func(i,j int) bool{
    // 	if squares[i][1] != squares[j][1]{
    // 		return squares[i][1] < squares[j][1]
    // 	}
    // 	return squares[i][1]+squares[i][2] < squares[i][1]+squares[i][2]
    // })
    // fmt.Println("total_area",total_area)

    // lb := func(target float64)int {
    // 	return sort.Search(n,func(i int)bool{
    // 		return float64(squares[i][1]+squares[i][2]) >= target
    // 	})
    // }
    is_dividing_area := func(target float64)bool{
    	// lower_bound,current_area := lb(target),float64(0)
    	current_area := 0.0
    	// fmt.Println("lb",lower_bound)
    	for _,sq := range squares{
    		// if float64(squares[lower_bound][1]) != target{
    		// 	y_len := min(float64(squares[lower_bound][1]+squares[lower_bound][2]),target)-float64(squares[lower_bound][1])
    		// 	current_area += float64(squares[lower_bound][2])*y_len
    		// 	fmt.Println("current_area in addition",lower_bound,current_area,y_len,squares[lower_bound][2])
    		// }
    		y_len := max(0.0,min(float64(sq[1]+sq[2]),target)-float64(sq[1]))
    		current_area += y_len*float64(sq[2])
    		fmt.Println("y len",y_len,sq[2],current_area)
    	}
    	fmt.Println("current",current_area)
    	return current_area >=total_area/2.0
    }

    start,end,eps := 0.0,float64(max_height),1e-5
    fmt.Println(max_height)
    for math.Abs(end-start)>eps{
    	mid := (start+end)/2.0
    	fmt.Println("mid",mid)
    	if is_dividing_area(mid){
    		end = mid
    	}else{
    		start = mid
    	}
    }
    return end
}


func main(){

	squares := [][]int{{0,0,1},{2,2,1}}
	fmt.Println(separateSquares(squares))
}