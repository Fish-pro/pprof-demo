package main

import "fmt"

func main(){
	/*
	空接口：不包含方法，所以任意类型都可以实现，因此可以存储任意类的数值
	*/
	var a1 A = Cat{"花猫"}
	var a2 A = Person{"yor",18}
	var a3 A = "haha"
	var a4 A = 100
	fmt.Println(a1, a2, a3, a4)

	testA(a1)
	testA(a2)
	testA(a3)
	testA(a4)
	testA(3.14)

	testA1(a1)
	testA1(a2)
	testA1(1.22)

	// map,key是字符串，value是任意类型
	map1 := make(map[string]interface{})
	map1["name"] = "york"
	map1["age"] = 19
	map1["friend"] = Person{"yang",18}
	fmt.Println(map1)

	// 切片，存储任意类型
	slice1 := make([]interface{},0,10)
	slice1 = append(slice1,a1,a2,a3,a4,1000)
	fmt.Println(slice1)
	test3(slice1)
}

// A是空接口，所以可以接受任意类型的参数
func testA(a A){
	fmt.Println(a)
}

func testA1(a interface{}){
	fmt.Println(a)
}

func test3(slice []interface{}){
	for i:=0;i<len(slice);i++{
		fmt.Printf("第%d个数据是：%v\n",i+1,slice[i])
	}
}

// 空接口
type A interface {

}

type Cat struct {
	Color string
}

type Person struct {
	Name string
	Age int
}