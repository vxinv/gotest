package main

import "fmt"

func main() {
	s0 := []int{1,2,3,4,222}
	fmt.Println("s0: ",s0,len(s0),cap(s0))
	var s1 []int = make([]int, 5, 8)
	var s2 []int = make([]int, 2, 6)
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := make([]int, 3, 7)
	fmt.Println(s3, len(s3), cap(s3))

	s4 := make([]int, 5, 8)
	for i := 0; i < len(s4); i++ {
		s4[i] = i + 1
	}
	var s5 = s4
	fmt.Println(s4, len(s4), cap(s4))
	fmt.Println(s5, len(s5), cap(s5))

	// 切片的遍历
	for index,value := range s5 {
		fmt.Println(index,value)
	}
	// 切片的追加
	ints := append(s5, 255)
	fmt.Println(ints)
	// 切割
	var s6 = s5[2:5]
	fmt.Println(s6)
	fmt.Println(s5)
	s5 = append(s5[:0], s5[1:]...)
	fmt.Println(s5)

	// 复制
	s7 := make([]int,5,8)
	for k:=0; k< len(s7); k++ {
		s7[k] = k+1;
	}
	fmt.Println(s7)
	var s8 = make([]int,2,6)
	copy(s8,s7)
	fmt.Println(s8)

}
