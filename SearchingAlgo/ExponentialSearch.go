package main 

import (
	"fmt"
	"math"
)

var index int
func ExponentialSearcher(SliceInt []int , ItemToSearch int)int{
	index=-1
	for i:=1 ; i < len(SliceInt) ; i*=2{
		if SliceInt[i] > ItemToSearch{
			for j:=i/2;j<i;j++{
				if SliceInt[j] == ItemToSearch{
					index=j
				}
			}
		}
	}

	for i:=int(math.Floor(math.Log(float64(len(SliceInt)))/math.Log(2.0))); i <len(SliceInt) ; i++{
		if SliceInt[i] == ItemToSearch{
			index=i
		}
	}

	return index
}


func main(){
	SliceForSeach:=[]int{10 ,22,23, 37, 54, 61, 89 ,234}
	ItemForSearch:=234
	ReturnedIndex:=ExponentialSearcher(SliceForSeach,ItemForSearch)
	fmt.Println(ReturnedIndex)

}