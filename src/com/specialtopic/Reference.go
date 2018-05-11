package main

import "fmt"

/**
  Go中关于值引用和参数引用demo：https://leokongwq.github.io/2017/01/22/golang-param-pass-value-or-point.html
 */
func main() {
	reference_test()
}

func reference_test()  {
	var m map[int]*int
	fmt.Println(&m)
	helper(m)
	fmt.Println(m == nil)
}
func helper(m map[int]*int)  {
	fmt.Println(&m)
	m = make(map[int]*int)
	//fmt.Println(m == nil)
}
