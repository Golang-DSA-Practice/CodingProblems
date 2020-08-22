package main

import (
	"fmt"
)


func QuickSorter(SliceInt []int,low int ,high int) []int {

	

	p:=SliceInt[0]

	for true{
		for i:=0;i<len(SliceInt);i++{
			if SliceInt[i] > p {
				low=i
				break
			}
		}

		for j:=len(SliceInt)-1;j>0;j--{
			if SliceInt[j] < p {
				high=j
				break
			}
		}

		if low < high {
			SliceInt[low],SliceInt[high]=SliceInt[high],SliceInt[low]

		}else {
			SliceInt[0],SliceInt[high]=SliceInt[high],SliceInt[0]
			
			QuickSorter(SliceInt[0:high-1],0,high-1)
			QuickSorter(SliceInt[high+1:],high+1,len(SliceInt)-1)
			break
			
			
		}
	}

	if SliceInt[low]==SliceInt[high]{
		return SliceInt
	}

	return make([]int,0,0)

	

}


func main(){
	SliceToSort:= []int{10,89,23,54,37,234,61,1,1002}
	SortedSlice:=QuickSorter(SliceToSort,0,len(SliceToSort)-1)
	fmt.Println(SortedSlice)

}