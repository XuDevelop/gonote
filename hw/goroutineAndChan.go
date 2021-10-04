package hw

import "fmt"

//func 判断一个数是否为素数
//func 分配任务

var primeChan chan int = make(chan int, 1)

func PrimeNum(i int) {
	//判断一个数是否为素数
	for j := 2; j < i; j++ {
		if i%j == 0 {
			return
		}
	}
	primeChan <- i

	//把结果放到管道里
}

func AssignTasks() {
	for i := 1; i <= 200000; i++ {
		go PrimeNum(i)
	}
	//把管道里的结果取出来
	var b bool
	for {
		select {
		case v := <-primeChan:
			fmt.Println(v, " ")
		default:
			b = true
		}
		if b {
			break
		}
	}
}
