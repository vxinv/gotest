package main

import "fmt"

func main()  {
	var s= "嘻哈china"
	for i := 1;i< len(s);i++ {
		fmt.Printf("%x \n",s[i])
	}
	// 按照字符打印
	for codeP,runeValue := range s{
		fmt.Printf("按照字符打印: %d ,%x \n",codeP,runeValue)
	}

	var s1 = "hello" // 静态字面量
	var s2 = ""
	for i:=0;i<10;i++ {
		s2 += s1 // 动态构造
	}
	fmt.Println("s1 len",len(s1))
	fmt.Println("s2 len",len(s2))

	var s3 = "hello world"
	var s4 = s3[3:8]
	fmt.Println(s4)

	var s5 = "hello world"
	var b = []byte(s5)  // 字符串转字节切片
	var s6 = string(b)  // 字节切片转字符串
	fmt.Println(b)
	fmt.Println(s6)
}
