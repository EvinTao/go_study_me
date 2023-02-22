// 指针
package main

import "fmt"

func main() {
	//准备一个字符串类型
	var house = "BeiJing no.10"

	//对字符串取地址，ptr 类型为 *string
	ptr := &house

	//打印ptr类型
	fmt.Printf("ptr type :%T \n", ptr)

	//打印ptr指针的地址
	fmt.Printf("address:%p \n", ptr)

	//对指针进行取值操作
	value := *ptr

	fmt.Printf("value type: %T \n", value)

	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)

}
