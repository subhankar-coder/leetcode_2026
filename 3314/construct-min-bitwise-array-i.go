package main

import(
	"fmt"
)


func minBitwiseArray(nums []int) []int {
    var ans []int 
    for _,val := range nums{
        flag := false
        for i:=1;i<=1000;i++{
            if i|(i+1) == val{
                ans = append(ans,i)
                flag=true
                break
            }
        }
        if !flag{
           ans = append(ans,-1)
        }
    }
    return ans
}

func main(){
	nums := []int{2,3,5,7}
	fmt.Println(minBitwiseArray(nums))
}