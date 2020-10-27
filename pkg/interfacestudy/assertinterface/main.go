package main

import (
	"fmt"
	"math"
)

func main(){

	var t1 Triangle = Triangle{3,4,5}
	fmt.Println(t1.peri(),t1.area())
	fmt.Println(t1.a,t1.b,t1.c)

	var c1 Circle = Circle{4}
	fmt.Println(c1.peri(),c1.area())
	fmt.Println(c1.radius)

	var s1 Shape = t1
	fmt.Println(s1.peri())
	fmt.Println(s1.area())

	var s2 Shape = c1
	fmt.Println(s2.peri())
	fmt.Println(s2.area())

	testShape(t1)
	testShape(c1)
	testShape(s1)

	getType(t1)
	getType(c1)
	getType(s1)

	var t2 *Triangle = &Triangle{2,3,4}
	fmt.Printf("t2:%T,%p,%p\n",t2,&t2,t2)
	getType(t2)

	getType1(t1)
	getType1(c1)
	getType1(s1)


}

func getType1(s Shape){
	switch ins:=s.(type) {
	case Triangle:
		fmt.Println("是三角形")
		fmt.Println(ins.a,ins.b,ins.c, ins.peri(),ins.area())
	case Circle:
		fmt.Println("是圆形")
		fmt.Println(ins.radius,ins.area(),ins.peri())
	default:
		fmt.Println("未知类型")
	}
}

func getType(s Shape){
	if ins,ok := s.(Triangle);ok{
		fmt.Println("是三角形")
		fmt.Println(ins.a,ins.b,ins.c, ins.peri(),ins.area())
	} else if ins,ok := s.(Circle);ok{
		fmt.Println("是圆形")
		fmt.Println(ins.radius,ins.area(),ins.peri())
	} else if ins,ok := s.(*Triangle);ok{
		fmt.Println("指针，三角形")
		fmt.Printf("t1:%T,%p,%p\n",ins,&ins,ins)
		fmt.Println(ins.a,ins.b,ins.c, ins.peri(),ins.area())

	}else {
		fmt.Println("未知类型")
	}
}

func testShape(sha Shape){
	fmt.Println("周长：",sha.peri())
	fmt.Println("面积：",sha.area())
}

// 1.定义接口
type Shape interface {
	peri() float64
	area() float64
}

// 2.定义实现类
type Triangle struct {
	//a float64
	//b float64
	//c float64
	a,b,c float64
}

func (t Triangle)peri()float64{
	return t.a+t.b+t.c
}

func (t Triangle)area()float64{
	p := t.peri()/2
	return math.Sqrt(p*(p-t.a)*(p-t.b)*(p-t.c))
}

type Circle struct {
	radius float64
}

func (c Circle)peri() float64{
	return 2*math.Pi*c.radius
}

func (c Circle)area()float64{
	return math.Pow(c.radius,2)*math.Pi
}