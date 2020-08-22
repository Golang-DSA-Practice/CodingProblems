package main



import (
    "fmt"
)

func CheckForSubstring(str,substr string) int{
	if len(substr) > len(str){
		return -1
	}else{
		first_char:=substr[0]
		
		
		for i:=0;i<len(str);i++{
			
			if str[i]==first_char{
				found:=0
				for j:=0;j<len(substr) && j<len(str)-i-1;j++{
					if substr[j]!=str[i+j]{
						found=1
						break
					}
				}
				if found==0{
					return i
				}
			}
		}
	}
	return -1

}
    
func main(){
	//find if a string is a substring of another string
	str:="I am writing a function to check if a string is substing of main string , returns -1 if not ,else returns the starting ind"
	substr:="index"
	fmt.Println(CheckForSubstring(str,substr))
}