package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
格式化打印占位符：
%v,原样输出
	%T，打印类型
	%t, bool类型
	%s，字符串
	%f，浮点
	%d，10进制的整数
	%b，2进制的整数
	%o，8进制
	%x，%X，16进制
		%x：0-9，a-f
		%X：0-9，A-F
	%c，打印字符
	%p，打印地址
	。。。
*/
func main() {
	const (
		FLAG bool = false
		num  int  = 13
	)
	// fmt.Printf 括号内必须是字符串形式打印
	// %T 表示打印类型
	// %v 表示打印值
	// 用逗号,隔开表示后面的数据 打印顺序
	// fmt.Printf("", FLAG, num)   // %!(EXTRA bool=false, int=13)
	// fmt.Printf("%T", FLAG, num) // bool%!(EXTRA int=13)
	// fmt.Printf("%v", FLAG, num) // false%!(EXTRA int=13)
	// fmt.Printf("%T,%v", FLAG, num) // bool,13

	// 省略var声明符
	// a := 100           //int
	// b := 3.14          //float64
	// c := true          // bool
	// d := "Hello World" //string
	// e := `Ruby`        //string
	// f := 'A'
	// fmt.Printf("%T,%b\n", a, a)       // int,1100100
	// fmt.Printf("%T,%f\n", b, b)       // float64,3.140000
	// fmt.Printf("%T,%t\n", c, c)       // bool,true
	// fmt.Printf("%T,%s\n", d, d)       // string,"Hello World"
	// fmt.Printf("%T,%s\n", e, e)       // string,Ruby
	// fmt.Printf("%T,%d,%c\n", f, f, f) // int32,65,A
	// fmt.Println("-----------------------")
	// fmt.Printf("%v\n", a) //
	// fmt.Printf("%v\n", b) //
	// fmt.Printf("%v\n", c) //
	// fmt.Printf("%v\n", d) //
	// fmt.Printf("%v\n", e) //
	// fmt.Printf("%v\n", f) //

	// var x int
	// var y float64
	// fmt.Println("请输入一个整数，一个浮点类型：")
	// fmt.Scanln(&x, &y) //读取键盘的输入，通过操作地址，赋值给x和y   阻塞式
	// fmt.Printf("x的数值：%d，y的数值：%f\n", x, y)
	// fmt.Scanf("%d,%f", &x, &y)
	// fmt.Printf("x:%d,y:%f\n", x, y)

	// bufio包中都是IO操作的方法：先创建Reader对象，然后读取
	fmt.Println("请输入一个字符串：")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n') // 会返回两个参数，所以s1, _ 的, _必须有，否则报错“..ReadString returns 2 values”
	fmt.Println("读到的数据：", s1)
}
