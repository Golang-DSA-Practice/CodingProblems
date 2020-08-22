package main 

import (
	"fmt"
)


func InsertionSorter(SliceInt []int) []int{
	for i:=1;i<len(SliceInt);i++{
		for j:=i;j>0;j--{
			if (SliceInt[j-1] > SliceInt[j]){
				SliceInt[j-1],SliceInt[j]=SliceInt[j],SliceInt[j-1]
			}
		}	
	}
	return SliceInt
}


func main(){
	SliceToSort := []int{10,89,23,54,37,234,61,1,1003}
	SortedSlice:=InsertionSorter(SliceToSort)
	fmt.Println(SortedSlice)

}