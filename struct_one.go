package main

import (
	"fmt"
)
func varDog() Dog {
	var duoduo Dog
	duoduo.Category = "pug"
	duoduo.Age = 3
	duoduo.Host = "孔得芳"

	fmt.Println("create with var", duoduo)
	return duoduo
}

func newDog() *Dog {
	duoduo := new(Dog)

	duoduo.Category = "pug"
	duoduo.Age = 3
	duoduo.Host = "孔得芳"

	fmt.Println("create with new:", duoduo)
	return duoduo
}

func dogFromAnon() Dog {
	duoduo := Dog{Category: "巴哥", Age: 3, Host: "芳大爷"}
	duoduo.Color = "灰色"
	fmt.Println("create with struct: ", duoduo)
	return duoduo
}

/**
 * package内部作用于，小写
 * package外部public，使用大写
 */
func (dog *Dog) Run() string {
	fmt.Println(dog, "Dog is running")
	return "Dog is Running"
}

func (dog *Dog) Eat() string {
	fmt.Println("Dog ---> yummy yummy!")
	return "Dog --> yummy yummy"
}

type Dog struct {
	Animal          // 组合，Dog里面包含Animal的所有字段和方法
	Category string
	Age      int
	Host     string
}


func (cat *Cat) Run() string {
	fmt.Println(cat, "Cat -->  is running")
	return "Cat -->  Running"
}

func (cat *Cat) Eat() string {
	fmt.Println("Cat --> yummy yummy!")
	return "Cat --> yummy yummy"
}

type Cat struct {
	Animal          // 组合，Dog里面包含Animal的所有字段和方法
	Category string
	Age      int
	Host     string
}

type Animal struct {
	Color string
}

