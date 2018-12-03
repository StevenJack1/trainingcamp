package main

import "fmt"

func main()  {


	a := 2
	b := 162
	var arr []int

	for i := 0; ; i++{            // 构造二叉树
		if i == 0 {
			arr = append(arr,a)
		} else {
			if 2 * arr[i-1] <= b {
				arr = append(arr,2 * arr[i-1])
			}
			if 10 * arr[i-1] + 1 <= b {
				arr = append(arr, 10 * arr[i-1] + 1)
			}
			if arr[i] == b {
				fmt.Println("Yes")
				break
			}
		}
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i],i)
	}
}



