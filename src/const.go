package main

import "fmt" // 格式化包

// 程序入口
func main() {
	// const 声明常亮
	// int 数字类型
	// string 字符串类型
	// bool 布尔类型
	const NUMBER int = 10
	const STRING string = "123"
	const a, b, c = 1, '1', true
	var area int
	area = NUMBER * a
	fmt.Printf("面积为：%d", area)
	println()
	println(a, b, c)

	// 常量可以作为枚举，常量组
	const (
		Zero = 0
		Number
		Male      = "男人"
		Female    = "女人"
		People    = "人类"
		ArePeople // 常量组中如不指定类型和初始化值，则与上一行非空常量右值相同
	)
	const (
		x uint16 = 16
		y
		s = "abc"
		z
	)
	// fmt.Printf("%T,%v\n", Number, Number)       // Number 会继承 Zero 的类型int以及值0
	// fmt.Printf("%T,%v\n", ArePeople, ArePeople) // ArePeople 会继承 People 的类型string以及值"人类"
	// fmt.Printf("%T,%v\n", y, y)                 // y 会继承 x 的类型uint16以及值16
	// fmt.Printf("%T,%v\n", z, z)                 // z 会继承 s 的类型string以及值"abc"

	// iota:特殊常量，可以认为是一个被编译器修改的常量
	// 第一个iota为0，每当iota在新的一行被使用的时候，它的值都会被自动+1.
	// 如果中断iota自增，则必须显式恢复。且后续自增值按行序递增
	// 自增默认是int类型，可以自行进行显示指定类型
	// 数字常量不会分配存储空间，无须像变量那样通过内存寻址来取值，因此无法获取地址
	const (
		a1 = iota // 0
		b1        // 1
		c1        // 2
		d1 = "ha" // 独立值，iota += 1
		e1        // "ha"   iota += 1
		f1 = 100  // iota +=1
		g1        // 100  iota +=1
		h1 = iota // 7,恢复计数
		i1        // 8
	)
	fmt.Println(a1, b1, c1, d1, e1, f1, g1, h1, i1)

	helloworld := "hello world"
	fmt.Println(helloworld)
}
