package main

import (
	"fmt"
	"math"
	"unsafe"
)

type Circle struct {
	x      int
	y      int
	Radius int
}

type ArrayStruct struct {
	value [10]int
}

type SliceStruct struct {
	value []int
}

func expandByValue(c Circle) {
	c.Radius *= 2
}

func expandByPointer(c *Circle) {
	c.Radius *= 2
}

// 面积
func (c Circle) Area() float64 {
	return math.Pi * float64(c.Radius) * float64(c.Radius)
}

// 周长
func (c Circle) Circumference() float64 {
	return 2 * math.Pi * float64(c.Radius)
}

// 扩大方法
func (c *Circle) expand() {
	c.Radius = c.Radius * 2
}
func (c Circle) expand1()  {
	c.Radius *= 2
}

type Point struct {
	x int
	y int
}

func (p Point) show() {
	fmt.Println(p.x , p.y)
}

type Circle1 struct {
	Point
	Radius int
}


func main() {
	c := Circle{
		x:      100,
		y:      100,
		Radius: 100,
	}
	fmt.Println("c: ", c)

	c1 := Circle{}
	fmt.Println(c1)

	var c2 *Circle = &Circle{100, 100, 50}
	fmt.Println("c2", c2)

	var as = ArrayStruct{[...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
	var ss = SliceStruct{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
	fmt.Println(unsafe.Sizeof(as), unsafe.Sizeof(ss))

	var c3 = Circle{Radius: 50}
	expandByValue(c3)
	fmt.Println(c3)
	expandByPointer(&c3)
	fmt.Println(c3)

	// 结构体方法
	var c4 = Circle{Radius: 50}
	fmt.Println(c4.Area(), c4.Circumference())
	// 指针变量调用方法形式上是一样的
	var pc = &c4
	fmt.Println(pc.Area(), pc.Circumference())

	c4.expand1()
	fmt.Println("c4: ",c4)
	c4.expand()
	fmt.Println("c4: ",c4)

	// Circle1
	var c5 = Circle1{
		Point : Point{
			x: 100,
			y: 100,
		},
		Radius: 50,
	}
	fmt.Printf("%+v\n", c5)
	fmt.Printf("%+v\n", c5.Point)
	fmt.Printf("%d %d\n", c5.x, c5.y) // 继承了字段
	fmt.Printf("%d %d\n", c5.Point.x, c5.Point.y)
	c5.show() // 继承了方法
	c5.Point.show()


}


