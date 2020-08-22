package main

import (
	"fmt"
)

func LinearSearcher(SliceInt []int , ItemToSearch int ) int {
	index:=-1 
	for i:=0;i<len(SliceInt);i++{
		if (SliceInt[i] == ItemToSearch){
			index=i
		}
	}
	return index	
}


func main(){
	SliceForSeach:=[]int{10,89,23,54,37,234,61,1,1002}
	ItemForSearch:=61
	ReturnedIndex:=LinearSearcher(SliceForSeach,ItemForSearch)
	fmt.Println(ReturnedIndex)

}



//Time Complexity O(n)
//Scans a unsorted array 