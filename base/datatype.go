package main

import (
	"bytes"
	"fmt"
	"strings"
)

func basename(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	index := strings.LastIndex(s, "/")
	s = s[index+1:]
	if dot := strings.Index(s, "."); dot > 0 {
		s = s[:dot]
	}
	return s
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func slice() {
	var x []int
	x = append(x, 1, 2, 3)
	x = append(x, 4, 5, 6)
	fmt.Printf("%#v \n", x)
}

func array1() {
	var a [3]int = [3]int{1, 2} // array of 3 integers
	fmt.Println(a[0])           // print the first element
	fmt.Println(a[len(a)-1])    // print the last element, a[2]
	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d -> %d\n", i, v)
	}
	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"
	for i, v := range q {
		fmt.Printf("%d -> %d\n", i, v)
	}
}

func make2new() {
	// new 函数
	p1 := new(int)
	fmt.Printf("p1 --> %#v \n ", p1)           //(*int)(0xc42000e250)
	fmt.Printf("p1 point to --> %#v \n ", *p1) //0

	var p2 *int
	i := 0
	p2 = &i
	fmt.Printf("p2 --> %#v \n ", p2)           //(*int)(0xc42000e278)
	fmt.Printf("p2 point to --> %#v \n ", *p2) //0

	//make() 函数
	var s1 []int
	if s1 == nil {
		fmt.Printf("s1 is nil --> %#v \n ", s1) // []int(nil)
	}
	s2 := make([]int, 3)
	if s2 == nil {
		fmt.Printf("s2 is nil --> %#v \n ", s2)
	} else {
		fmt.Printf("s2 is not nill --> %#v \n ", s2) // []int{0, 0, 0}
	}

	// map
	var m1 map[int]string
	if m1 == nil {
		fmt.Printf("m1 is nil --> %#v \n ", m1) //map[int]string(nil)
	}

	m2 := make(map[int]string)
	if m2 == nil {
		fmt.Printf("m2 is nil --> %#v \n ", m2)
	} else {
		fmt.Printf("m2 is not nill --> %#v \n ", m2) //map[int]string{}
	}

	// chan
	var c1 chan string
	if c1 == nil {
		fmt.Printf("c1 is nil --> %#v \n ", c1) //(chan string)(nil)
	}

	c2 := make(chan string)
	if c2 == nil {
		fmt.Printf("c2 is nil --> %#v \n ", c2)
	} else {
		fmt.Printf("c2 is not nill --> %#v \n ", c2) //(chan string)(0xc420016120)
	}

}

func modifyMap(m map[int]string) {
	m[0] = "string"
}

func modifyChan(c chan string) {
	c <- "string"
}

func modify() {
	m2 := make(map[int]string)
	if m2 == nil {
		fmt.Printf("m2 is nil --> %#v \n ", m2)
	} else {
		fmt.Printf("m2 is not nill --> %#v \n ", m2) //map[int]string{}
	}

	modifyMap(m2)
	fmt.Printf("m2 is not nill --> %#v \n ", m2) // map[int]string{0:"string"}

	c2 := make(chan string)
	if c2 == nil {
		fmt.Printf("c2 is nil --> %#v \n ", c2)
	} else {
		fmt.Printf("c2 is not nill --> %#v \n ", c2)
	}

	go modifyChan(c2)
	go modifyChan(c2)
	fmt.Printf("c2 is not nill --> %#v ", <-c2) //"string"
}

type Foo struct {
	name string
	age  int
}

func structDemo() {
	//声明初始化
	var foo1 Foo
	fmt.Printf("foo1 --> %v\n ", foo1)
	fmt.Printf("foo1 --> %+v\n ", foo1)
	fmt.Printf("foo1 --> %#v\n ", foo1)

	foo1.age = 1
	fmt.Println(foo1.age)

	//struct literal 初始化
	foo2 := Foo{}
	fmt.Printf("foo2 --> %+v\n ", foo2)
	foo2.age = 2
	fmt.Println(foo2.age)

	//指针初始化
	foo3 := &Foo{}
	fmt.Printf("foo3 --> %+v\n ", foo3)
	foo3.age = 3
	fmt.Println(foo3.age)

	//new 的初始化
	foo4 := new(Foo)
	fmt.Printf("foo4 --> %+v\n ", foo4)
	foo4.age = 4
	fmt.Println(foo4.age)

	//声明指针并用 new 初始化
	var foo5 *Foo = new(Foo)
	fmt.Printf("foo5 --> %+v\n ", foo5)
	foo5.age = 5
	fmt.Println(foo5.age)
}

func main() {
	//println(basename("/home/taoyf/zhangsan.img"))
	//println(basename2("/home/23233/343434.txt"))
	//println(comma("1234567890"))
	//fmt.Println(intsToString([]int{1, 2, 3}))

	//x := 123
	//y := fmt.Sprintf("%d", x)
	//fmt.Println(y, strconv.Itoa(x))
	//fmt.Println(strconv.FormatInt(int64(x), 2))

	//array1()
	//make2new()
	//modify()
	//slice()
	structDemo()
}
