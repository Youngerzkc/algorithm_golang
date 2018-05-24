package main

import(
	"fmt"
)
func main(){
	var aaa =[]int{1,2,3,5}
	var bbb [4]int
	size:=len(aaa)
	fmt.Println("size is ",size)
	for k,v:=range aaa{
		if k ==size-1{
			bbb[k]=v+1
			break
		}
		bbb[k]=aaa[k]
		fmt.Println("k is ",k,v)
	}
	fmt.Println("aaa is ",aaa)
	fmt.Println("bbb is ",bbb)

}