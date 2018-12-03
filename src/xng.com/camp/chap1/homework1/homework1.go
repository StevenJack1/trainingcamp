package main

import "fmt"

func main()  {
	var arr = []int {1,2,3}
	fmt.Println(Homework1(arr))
}
func Homework1(arr []int)  int{
	max := 0
	for i := 0; i < len(arr);  i++{
		for j := i+1; j < len(arr); j++ {
			if (arr[i] - i) == (arr[j] - j)  {
				if j-i > max {
					if arr[j] == 1000 || arr[i] == 1{
						max = j -i
					} else {
						max = j - i -1
					}
				}
			}
		}
	}
	return max

}