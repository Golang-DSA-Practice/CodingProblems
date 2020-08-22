package main

import (
	"fmt"
)

func BubbleSorter(SliceInt []int) []int{ 
	isSwap:=true

	for isSwap {
		isSwap=false
		for ele:=0 ; ele < (len(SliceInt)-1); ele++ { 
			if (SliceInt[ele] < SliceInt[ele + 1]){ 
				temp:=SliceInt[ele]
				SliceInt[ele] = SliceInt[ele + 1]
				SliceInt[ele+1]=temp
				isSwap=true
			}
		}
	}

	return SliceInt
}

func main(){
	SliceToSort := []int{10,89,23,54,37,234,61}
	SortedSlice:=BubbleSorter(SliceToSort)
	fmt.Println(SortedSlice)

}

//time complexity O(n^2)