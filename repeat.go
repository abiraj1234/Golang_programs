package main

import "fmt"

func findRepeatedElement(nums []int) int {
    
    count := make(map[int]int)
    for _, num := range nums {
        count[num]++
        
        if count[num] == len(nums)/2 {
            return num
        }
    }
 
    return -1
}

func main() {
    nums := []int{1, 2, 3, 3, 4, 5, 6}
    repeated := findRepeatedElement(nums)
    fmt.Println(repeated)
}