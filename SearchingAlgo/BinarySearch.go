package main 

import (
	"fmt"
)
var index int
func BinarySearcher(SliceInt []int,ItemToSearch int, Front,Rear int) int {
	index=-1
	
	if Rear >= Front{	
		Mid:=(Rear+Front)/2
		if (SliceInt[Mid] == ItemToSearch){
			
			index=Mid
			
		}else if (SliceInt[Mid] > ItemToSearch){
			
			BinarySearcher(SliceInt,ItemToSearch,Front,Mid-1)
			
		}else{
			
			BinarySearcher(SliceInt,ItemToSearch,Mid+1,Rear)
			
		}
	}
	

	return index
}



func main(){
	SliceForSeach:=[]int{10 ,23, 37, 54, 61, 89 ,234}
	ItemForSearch:=89
	ReturnedIndex:=BinarySearcher(SliceForSeach,ItemForSearch,0,len(SliceForSeach)-1)
	fmt.Println(ReturnedIndex)

}

//O(logn)