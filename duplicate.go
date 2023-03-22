package main 


import "fmt"


func containsDuplicate(nums []int) bool {

    temp := make(map[int]int)
	
	for _, dup:= range nums {
	   _, con:= temp[dup]
	   
	   if con{
		
		return true 
		}else {
		
		temp[dup] = 0 
		}
	 
	}
	return false
	
}

func main() {
	nums:= []int{1,2,3,1}
	fmt.Println("the duplicate value is :",  containsDuplicate(nums))
}