package main 

import (
	"fmt"
)


func SelectionSorter(SliceInt []int) []int{
	for i:=0 ; i<len(SliceInt) ;i++{
		min:=SliceInt[i]
		
		for j:=i+1;j<len(SliceInt) ;j++ {
			if (SliceInt[j] < min){
				min=SliceInt[j]
				temp:=SliceInt[i]
				SliceInt[i]=SliceInt[j]
				SliceInt[j]=temp
			}

		}
		

	}

	return SliceInt

}


func main(){
	SliceToSort := []int{10,89,23,54,37,234,61}
	SortedSlice:=SelectionSorter(SliceToSort)
	fmt.Println(SortedSlice)

}


