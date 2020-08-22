package main

import (
	"fmt"
	"math"
)

var index int
func JumpSearcher(SliceInt []int,ItemToSearch int) int{
	index:=-1
	JumpStep:= math.Sqrt(float64(len(SliceInt))) // max scans n/m + m-1 = 0 , you will get m=sqrt(m)
	
	for i:=0;i<len(SliceInt);i+=int(JumpStep){
		if (SliceInt[i] > ItemToSearch){
			for j:=i-int(JumpStep) ; j < i ; j++{
				if (SliceInt[j]==ItemToSearch){
					index=j
				}
			}
		}
	}
	return index
}


func main(){
	SliceForSeach:=[]int{10 ,23, 37, 54, 61, 89 ,234}
	ItemForSearch:=60
	ReturnedIndex:=JumpSearcher(SliceForSeach,ItemForSearch)
	fmt.Println(ReturnedIndex)

}



//O(sqrt(n))