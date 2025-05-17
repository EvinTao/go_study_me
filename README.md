# go_study_me

## base
iota 类似行所以，每次新增一行都会+1;


基础
## 变量
格式：var name type = expression   
类型和表达式部分只能省略一个，不能都省略；如果表达式省略，其初始值对应于类型的零值。
对于int ~0, bool ~ false, string ~"",接口和引用(slice,指针，map,通知，函数) ~ nil ;
零值机制保障所有的变量是良好定义，GO里面不存在未初始化的变量。

## 短变量
name :=  expression
短变量声明主要用于局部变量的声明和初始化中


## 指针
指针的值是一个变量的地址
取地址的操作符 &


## 生命周期
包级别的变量的生命周期是整个程序的执行时间。
局部变量是动态的生命周期，变量一直生存到它变得不可访问。


## 类型声明
格式：type name underlying-type   --> 可以理解为别名
例如：type Celsius float64  
goroutine