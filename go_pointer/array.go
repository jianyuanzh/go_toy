package go_pointer

import "fmt"

//指针数组: 数组里面装的是指针
//数组指针：指针指向一个数组

func TestPointerArray() {
	a, b := 1, 2
	pointArr := [...]*int {&a, &b}
	fmt.Println("指针数组：", pointArr)

//	数组指针
	arr := [...]int {1, 2, 3, 4, 6}
	arrPointer := &arr
	fmt.Println("数组指针：", arrPointer)
}
