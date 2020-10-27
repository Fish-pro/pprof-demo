package main

import "fmt"

func main(){
	var cat Cat = Cat{}
	cat.test1()
	cat.test2()
	cat.test3()

	fmt.Println("------------------")
	var a1 A = Cat{}
	a1.test1()

	var b1 B = Cat{}
	b1.test2()

	var c1 C = Cat{}
	c1.test1()
	c1.test2()
	c1.test3()

	fmt.Println("------------------")
	//var c2 C = a1 error:C接口有自己定义的方法
	var a2 A = c1
	a2.test1()
}

type A interface {
	test1()
}

type B interface {
	test2()
}

type C interface {
	A
	B
	test3()
}

type Cat struct { // 如果想实现接口C，那么就要把接口A和B中的方法也实现

}

func (c Cat)test1(){
	fmt.Println("test1")
}

func (c Cat)test2(){
	fmt.Println("test2")
}

func (c Cat)test3(){
	fmt.Println("test3")
}