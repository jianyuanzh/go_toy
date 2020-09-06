package go_pointer

import "fmt"

func ShowPointer() {
	var value int = 100
	var pointer *int = &value

	fmt.Println("普通值：", value)
	fmt.Printf("地址:0x%x\n", &value)
	fmt.Printf("指针指向的地址: 0x%x\n", pointer)
	fmt.Printf("指针指向的地址对应的值：%d\n", *pointer)
}
