package main

import (
	f "fmt"
	"sync"
	"time"
)

// 常量
const a string = "abc"
const b = "abc"

func main() {
	f.Println(a, b)
	gogoMux()
	chans()
	selects()
}

/*func arr() {
	var arr [4]int
	var arr2 = [2]int{2, 3}
	var arr3 = [...]int{2, 5}

	var arr4 []int  // 切片
	var arr5 [2]int // 数组

	arr5 = append(arr4, 1) // 不合法
	arr5 = append(arr5, 1) // 不合法
	arr5 = arr2            // 合法
	arr5 = arr3            // 合法
}*/

func maps() {
	var m map[string]string
	m["test"] = "333" // 不合法，编译报错，m未初始化

	m = make(map[string]string, 10) // 初始化，且起始容量为10
	m2 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	// 删除元素
	delete(m2, "key1")

	// 判断key是否存在
	if v, ok := m2["key1"]; !ok {
		f.Println("no such key")
	} else {
		f.Println(v)
	}

	// 不可以，零值特性
	if m2["key1"] != "" {
		f.Println("no such key")
	}
}

func flowControl() {
	if a == "abc" {
		f.Println(a)
	}

	a := "abc"
	switch a {
	case "abc":
		f.Println(1)
	case "cba":
		f.Println(2)
		fallthrough
	case "bca":
		f.Println(3)
	}

	typeSwitch(21)
	typeSwitch("hello")
	typeSwitch(true)
}

func typeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		f.Printf("Twice %v is %v\n", v, v*2)
	case string:
		f.Printf("%q is %v bytes long\n", v, len(v))
	default:
		f.Printf("I don't know about type %T!\n", v)
	}
}

// 函数直接演示PPT上的例子

// 结构体和接口一起演示

type Human struct {
	Name string
}

func structs() {
	var h *Human // 引用（指针）类型
	var h1 Human // 值类型

	h.Name = "张三" // 非法，引用类型要初始化，不适用零值特性
	h1.Name = "张三"

	h = &Human{}   // 初始化张三
	h = new(Human) // 初始化张三，二者效果一样
	h = &Human{
		Name: "张三",
	}

	h = &Human{
		"张三",
	}
}

type animal interface {
	Name()
	Eat(food string)
	Run()
}

type Dog struct {
}

// 并发协程
func GoRun(i int) {
	f.Println(i)
}

func gogo() {
	// 使用go来发起协程
	go GoRun(1)
	go GoRun(2)

	time.Sleep(1 * time.Second)
}

func gogoMux() {
	// 使用go来发起协程
	w := sync.WaitGroup{}

	w.Add(1)
	go func(n int) {
		f.Println(n)
		w.Done()
	}(1)

	w.Add(1)
	go func(n int) {
		f.Println(n)
		w.Done()
	}(1)

	w.Wait()
}

// 并发chan
func sendData(ch chan int) {
	ch <- 1
	ch <- 2
}

func getData(ch chan int) {
	for {
		i := <-ch
		f.Println(i)
	}
}

func chans() {
	ch := make(chan int)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1 * time.Second)
}

func getData2(ch1 chan int, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			f.Printf("got data from ch1: %d\n", v)
		case v := <-ch2:
			f.Printf("got data from ch2: %d\n", v)
		}
	}
}

func selects() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go sendData(ch1)
	go sendData(ch2)
	go getData2(ch1, ch2)

	time.Sleep(1 * time.Second)
}

// 单元测试，直接使用PPT示例
