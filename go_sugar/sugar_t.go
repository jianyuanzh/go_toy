package go_sugar

import "fmt"

func Sugar(vals ... string) {
	for _,v := range vals {
		fmt.Println("----> v:", v)
	}
}


func Sugar2() {
	value := "A"
	fmt.Println("----> v:", value)
	value = "b"
	fmt.Println("----> v:", value)
}


