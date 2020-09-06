package main

import "fmt"


// len -> string, slice, array, map, chan
// cap -> slice, array, chan
// close -> chan
func lenSlice() {
	mSlice := make([]string, 2, 5)
	fmt.Println("len:", len(mSlice))
	fmt.Println("cap:", cap(mSlice))

	mSlice[0] = "string"
	mSlice[1] = "string"
	fmt.Println("len:", len(mSlice))
	fmt.Println("cap:", cap(mSlice))

	mSlice = append(mSlice, "code")
	fmt.Println("len:", len(mSlice))
	fmt.Println("cap:", cap(mSlice))

	mSlice = append(mSlice, "code")
	mSlice = append(mSlice, "code")
	mSlice = append(mSlice, "code")
	mSlice = append(mSlice, "code")
	fmt.Println("len:", len(mSlice))
	fmt.Println("cap:", cap(mSlice))

	//mSlice[2] = "string"
}