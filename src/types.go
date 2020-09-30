package main

// 格式化包

/**
 数值型
1、整数型
int8 有符号 8 位整型 (-128 到 127) 长度：8bit
int16 有符号 16 位整型 (-32768 到 32767)
int32 有符号 32 位整型 (-2147483648 到 2147483647)
int64 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
uint8 无符号 8 位整型 (0 到 255) 8位都用于表示数值：
uint16 无符号 16 位整型 (0 到 65535)
uint32 无符号 32 位整型 (0 到 4294967295)
uint64 无符号 64 位整型 (0 到 18446744073709551615)

2、浮点型
float32
IEEE-754 32位浮点型数
float64
IEEE-754 64位浮点型数
complex64
32 位实数和虚数
complex128
64 位实数和虚数
3、其他
byte
类似 uint8
rune
类似 int32
uint
32 或 64 位
int
与 uint 一样大小
uintptr
无符号整型，用于存放一个指针
1.3 字符串型 string
1.4 数据类型转换：Type Convert
语法格式：Type(Value)
常数：在有需要的时候，会自动转型
变量：需要手动转型 T(V)
注意点：兼容类型可以转换

复合类型(派生类型):
1、指针类型（Pointer）
2、数组类型
3、结构化类型(struct)
4、Channel 类型
5、函数类型
6、切片类型
7、接口类型（interface）
8、Map 类型
*/

func main() {
	const (
		FLAG bool = false
		num  int  = 13
	)
	println(FLAG, num) // bool,13
}
