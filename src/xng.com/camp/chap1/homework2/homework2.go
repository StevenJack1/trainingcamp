package main

import "fmt"

func main()  {
	//var n = 3
	//var m = 6
	var a = [3]int {5,10000000000}
	var b = []int {5,6,9999999999}      // key
	var temp [3]int
	for i := 0; i < len(a); i++ {
		if i == 0 {
			temp[i] = a[i]
		} else {
			temp[i] = temp[i-1] + a[i]
		}
	}   // temp[] = { 5,1000000005}
	for i := 0; i < len(b); i++ {          // 求宿舍楼号和房间号
		var homeNum int
		dormitoryNum := binarySearch(temp,0,len(temp),b[i])
		if dormitoryNum == 1 {
			homeNum = b[i]
		} else{
			homeNum = b[i] - temp[dormitoryNum - 2]          // 计算当前这个宿舍口号的房间号码
		}

		fmt.Println(dormitoryNum,homeNum)
	}

}

func binarySearch(a [3]int,low int,high int,key int)  int{   // low   high  key=5
	mid := (high - low)/2 + low
	if low <= high {
		if key == a[mid] {
			return mid+1
		} else if key > a[mid] {
			return binarySearch(a,mid + 1,high,key)
		} else if key < a[mid] {
			return binarySearch(a,low,mid - 1,key)
		}
	}
	return low+1

}

