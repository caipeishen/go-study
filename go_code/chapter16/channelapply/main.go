package main

import (
	"fmt"
	"time"
)

//write Data
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		//放入数据
		//time.Sleep(time.Second * 5)
		intChan <- i //
		fmt.Println("writeData ", i)
		//time.Sleep(time.Second)
	}

	fmt.Println("关闭iniChan")
	close(intChan) //关闭
}

//read data
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		time.Sleep(time.Second)
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	//readData 读取完数据后，即任务完成
	exitChan <- true
	close(exitChan)

}

func simple() {
	//创建两个管道
	intChan := make(chan int)
	//exitChan := make(chan bool, 1)

	intChan <- 1
	go func() {
		<-intChan
	}()
	fmt.Printf("intChan:")
}

func closeDetail() {
	//创建两个管道
	intChan := make(chan int, 1)
	//exitChan := make(chan bool, 1)

	intChan <- 1

	close(intChan)

	num, ok := <-intChan
	fmt.Printf("num=%v, ok=%v\n", num, ok)

	num2, ok2 := <-intChan
	fmt.Printf("num2=%v, ok2=%v\n", num2, ok2)
}

func main() {
	//simple()

	//closeDetail()

	//创建两个管道
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	//time.Sleep(time.Second)
	for {

		// 为什么不能使用len()==0判断推出，因为len代表缓冲区的数据长度，开始的时候len()为0，其次每次把缓冲区的数据取完，len也就为0
		//if len(exitChan) == 0 {
		//	break
		//}
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}
