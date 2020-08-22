package main 

import (
	"fmt"
)

var index int 

func InterpolationSearcher(SliceInt []int,ItemToSearh,low,high int) int {
	/*
		Interpolation search is an improvement over Binary 
		Used on sorted uniformly distributes array
		does not go to the middle of an array to seach for the element 
		instead goes to a POS defined by below formula

		pos=low + [(x-arr[low])*(high-low)/(arr[high])-arr[low]]
		high>=low is not enough as  x -arr[low] will result into negative value and x = arr[high] will result into [(x-arr[low])*(high-low)/(arr[high])-arr[low]] 
		low + high -low= high 

		and x > arr[high] will result into something greater than high because x-arr[low]/arr[high]-arr[low] will be greater than 1 

		normal maths suggests x > arr[low] and x< arr[high]
	*/
	index=-1
	pos:=(ItemToSearh - SliceInt[low])*(high-low)/(SliceInt[high]-SliceInt[low])
	pos=pos + low


	if (high>=low && ItemToSearh > SliceInt[low] && ItemToSearh < SliceInt[high]) {
		if (SliceInt[pos] == ItemToSearh){
			index=pos
		}else if (SliceInt[pos] < ItemToSearh){
			InterpolationSearcher(SliceInt,ItemToSearh,pos + 1,high)
		}else {
			InterpolationSearcher(SliceInt,ItemToSearh,low,pos-1)
		}
	}
	return index
}

func main(){
	SliceForSeach:=[]int{10 ,23, 37, 54, 61, 89 ,234}
	ItemForSearch:=0
	ReturnedIndex:=InterpolationSearcher(SliceForSeach,ItemForSearch,0,len(SliceForSeach)-1)
	fmt.Println(ReturnedIndex)

}