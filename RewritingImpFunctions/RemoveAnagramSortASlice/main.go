package main



import (
    "fmt"
)


func CompareStr(str1,str2 string) bool {
	if len(str1)==len(str2){
		for _,v1:=range str1{
			match:=false
			for _,v2:=range str2{
				if v1==v2{
					match=true
					break
				}
			}
			if match == false{
				return false
			}
		}
	}else{
		return false
	}
	return true
}

func Contains(slc []string,val string) bool {
	for _,v :=range slc{
		if CompareStr(v,val){
			return true
		}
	}
	return false
}


func SortStringInArray(slc []string) []string{
	exchange:=true
	for exchange == true{
		exchange=false
		for i:=0;i<len(slc)-1;i++{
			if slc[i] < slc[i+1]{
					slc[i],slc[i+1]=slc[i+1],slc[i]
					//fmt.Println(slc)
					exchange=true
					
			}
		}
		

		
	}
	return slc
}

	

func main() {
	var nonDupArray []string
	str:=[]string{"code","doce","ecod","framer","frmare","hello","world","piyush"}
	for i:=0;i<len(str);i++{
		for j:=i+1;j<len(str);j++{
			if CompareStr(str[i],str[j]){
				
				if  ! Contains(nonDupArray,str[i]){
					nonDupArray=append(nonDupArray,str[i])
				}
				
			}else{
				if  ! Contains(nonDupArray,str[j]){
					nonDupArray=append(nonDupArray,str[j])
				}
			}

		}
	}
	
	fmt.Println(SortStringInArray(nonDupArray))
	
    

}