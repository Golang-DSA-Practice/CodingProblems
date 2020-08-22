package main 

import (
	"fmt"
)

var found string
func SublistSearcher(BiggerSliceInt,SmallerSliceInt []int) string{
	found="No"
	for i:=0;i<len(BiggerSliceInt);i++{
	
		if (BiggerSliceInt[i]==SmallerSliceInt[0]){
			found="Yes"
			for j:=0;j<len(SmallerSliceInt);j++{
				if (j+i < len(BiggerSliceInt)){
					if (SmallerSliceInt[j] != BiggerSliceInt[j+i]){
						found="No"
						break
					}
				}else{
					found="No"
					break
				}
			}
			
		}
		if (found=="Yes"){
			break
		}
	}
	return found
}



func main(){

	BiggerSliceInt:=[]int{1,2,3,1,2,1,2,3,4,5}
	SmallerSliceInt:=[]int{1,10,12}
	Found:=SublistSearcher(BiggerSliceInt,SmallerSliceInt)
	fmt.Println(Found)
}