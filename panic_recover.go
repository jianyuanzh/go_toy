package main

import (
	"fmt"
)

// len -> string, array, slice, map,
// cap ->
//func main() {
//	//defer func() {
//	//	message := recover()
//	//	fmt.Println("panic message: ", message)
//	//}()
//}


// test panic revoer
func receivePanic() {
	//panic("This is a panic")
	//panic(errors.New("this is error "))
	panic(123)
}

func recoverPanic() {
	message := recover()
	fmt.Println("panic: ",  message)

	switch message.(type) {
	case string:
		fmt.Println("string message:", message)
	case error:
		fmt.Println("error message:", message)
	default:
		fmt.Println("unknown message:", message)
	}
}