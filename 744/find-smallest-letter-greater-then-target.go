package main

import (
	"fmt"
	"sort"
)


func nextGreatestLetter(letters []byte, target byte) byte {
    n := len(letters)
	if target > letters[n-1]{
		return letters[0]
	}
	
	ele := sort.Search(n,func(i int) bool {
		return letters[i] > target
	})
	// fmt.Println(ele)
	return byte(letters[min(ele,n-1)])
}

func main(){
	
	letters,target := []byte{'c','f','j'},'j'
	fmt.Println(nextGreatestLetter(letters,byte(target)))
}