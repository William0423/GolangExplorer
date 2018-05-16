package main

import (
	"fmt"
	"time"
)

/**
example--1
 */
//func print() {
//	fmt.Println("Hello world")
//}
//func main() {
//	for i :=0 ; i < 10; i++  {
//		go print()
//	}
//}

/**
example--2
 */
//func print(ch chan int) {
//	fmt.Println("Hello world")
//	ch<- 1
//}
//func main() {
//	chs := make([]chan int, 10)
//	for i := 0; i < 10; i++ {
//		chs[i] = make(chan int)
//		go print(chs[i])
//	}
//	for _, ch := range(chs){
//		<- ch // 相当于 ： value := <- ch
//	}
//}


/**
example--3
http://colobu.com/2016/04/14/Golang-Channels/
 */
//func main() {
//	go func() {
//		time.Sleep(1 * time.Hour)
//	}()
//	c := make(chan int)
//	go func() {
//		for i := 0; i < 10; i = i + 1 {
//			c <- i
//		}
//		close(c)
//	}()
//	for i := range c {
//		fmt.Println(i)
//	}
//	fmt.Println("Finished")
//}

/**
example--4
http://colobu.com/2016/04/14/Golang-Channels/
 */
//func main() {
//	c1 := make(chan string, 1)
//	go func() {
//		time.Sleep(time.Second * 2)
//		c1 <- "result 1"
//	}()
//	select {
//	case res := <-c1:
//		fmt.Println("=============")
//		fmt.Println(res)
//
//		fmt.Println("###########")
//	case <-time.After(time.Second * 3):
//		fmt.Println("@@@@@@@@@@@@@@")
//		fmt.Println("timeout 1")
//	}
//}


const StudentNum = 30
type HomeWork struct {
	Context string
}
//func student(hwChan chan HomeWork) {
func student(hwChan chan <- HomeWork) { // 和上一行不同的是，此处表示只能向这个channel写数据
	//学生提交作业
	var h  HomeWork
	h.Context = "hello"
	hwChan <- h
}
//func teacher(hwChan chan HomeWork){
func teacher(hwChan <-chan HomeWork) { // 表示只能从这个channel中读数据

//老师取出30个学生的作业批改c
	for  i:=0;i<StudentNum;i++ {
		hw := <- hwChan
		fmt.Println(hw)
	}
}

func main(){
	hwChan := make(chan HomeWork,30)
	//每个学生开启一个goroutine，学生单独做作业，做完作业提交到hwChan中即可
	for i:=0;i<StudentNum;i++ {
		go student(hwChan)
	}
	//老师开启一个goroutine，去批改学生作业
	go teacher(hwChan)

	time.Sleep(5*time.Second) // 此处可以替换为一个chan
}