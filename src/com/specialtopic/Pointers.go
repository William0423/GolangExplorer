package main

import "fmt"


type Student struct {
	name  string
	id    int
	score uint
}

func main()  {
	//pointers_null()
	reference_test()
	// memery_test()
}

/**
	指针是指向某个变量内存地址的值: 取变量a的内存地址，赋给指针b，然后对指针b做操作，此时操作的是a对应内存空间的值
	https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back
 */
func pointers_definition()  {
	a := 200
	b := &a
	*b++
	fmt.Println(a)
}


func pointers_null()  {
	stu1 := &Student{}
	fmt.Println(stu1.name)
	if stu1 == nil {
		fmt.Printf("yes")
	} else {
		fmt.Println("no")
	}

	// 定义一个指针变量，不指向任何内存空间的地址
	var stu2 *Student
	//stu2 = stu1 // 把stu1的地址赋值给stu2
	if stu2 == nil {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}


}

func memery_test() {
	//new分配出来的数据是指针形式
	p := new(Student)
	p.name = "China"
	p.id = 63333
	p.score = 99
	fmt.Println(*p)
	//var定义的变量是数值形式
	var st Student
	st.name = "Chinese"
	st.id = 666333
	st.score = 100
	fmt.Println(st)

	//make分配slice、map和channel的空间，并且返回的不是指针
	var ptr *[]Student
	fmt.Println(ptr)     //ptr == nil

	ptr = new([]Student) //指向一个空的slice
	fmt.Println(ptr)

	*ptr = make([]Student, 3, 100)
	fmt.Println(ptr)

	stu := []Student{{"China", 3333, 66}, {"Chinese", 4444, 77}, {"Chince", 5555, 88}}
	fmt.Println(stu)
}