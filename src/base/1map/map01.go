package main

import (
	"fmt"
)

func main()  {
	m1 := make(map[int]string)
	fmt.Println(m1)
	m2 := map[int]string{
		90: "优秀",
		80: "良好",
		60: "及格",
	}
	fmt.Println(m2)
	m3 := make(map[int]string,16)
	fmt.Println(m3)

	m4 := map[string]int{
		"apple":  2,
		"banana": 3,
		"orange": 4,
	}
	names := make([]string,0,3)
	scores := make([]int,0,3);
	for key,value := range m4 {
		names = append(names, key)
		scores = append(scores, value)
	}
	fmt.Println(names)
	fmt.Println(scores)

}
