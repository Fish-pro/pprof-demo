package main

import "fmt"

func main(){
	/*
	1.当需要接口需要对象时，可以使用任意类型对象代替
	2.接口对象不能访问对象类中的属性
	3.多态：一个事物的多种形态
		go通过接口模拟多态
			1.看成实现类结构体，可以访问结构体中的属性和方法
			2.看成接口类，只能访问接口中的方法
	4.接口的使用：
		1.如果一个函数一接口类型作为接受参数，那么可以传入实现该接口方法的任意实现结构体对象
		2.定义一个接口类对象，那么可以赋值任意实现该接口的结构体对象
	5.鸭子类型：有鸭子的行为就可以看成是鸭子，关注接口实现的方法，而不是类型本身
	*/
	m1 := Muse{"鼠标"}
	fmt.Println(m1.Name)

	f1 := FlashDish{"u盘"}
	fmt.Println(f1.Name)

	testInterface(m1)
	testInterface(f1)

	var usb USB
	usb = m1
	usb.start()
	usb.end()

	f1.delete()

	var arr [3]USB
	arr[0]=m1
	arr[1]=f1
	fmt.Println(arr)
}

// 1.定义接口
type USB interface {
	start()
	end()
}

// 2.实现类
type Muse struct {
	Name string
}

type FlashDish struct {
	Name string
}

func (m Muse)start(){
	fmt.Println(m.Name,"已就绪，可以开始工作")
}

func (m Muse)end(){
	fmt.Println(m.Name,"结束工作，可以安全退出")
}

func (f FlashDish)start(){
	fmt.Println(f.Name,"准备开始工作")
}

func (f FlashDish)end(){
	fmt.Println(f.Name,"结束工作，可以拔出")
}

// 测试接口方法
func testInterface(usb USB){
	usb.start()
	usb.end()
}

func (f FlashDish)delete(){
	fmt.Println(f.Name,"删除了数据")
}