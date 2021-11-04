package main

import (
	"errors"
	"fmt"
)

type StackStruct struct {
	Cap int //总共容量
	Top int
	Arr []int
}

func (s *StackStruct) Push(v int) (err error) {
	if s.Top == s.Cap-1 {
		return errors.New("stack full")
	}
	s.Top++
	s.Arr[s.Top] = v
	return
}

func (s *StackStruct) Pop() (v int, err error) {
	if s.Top == -1 {
		return 0, errors.New("stack empty")
	}
	v = s.Arr[s.Top]
	s.Top--
	return
}

//运算的方法
func Cal(num1 int, num2 int, oper int) int {
	switch oper {
	case '+':
		return num1 + num2
	case '-':
		return num1 - num2
	case '*':
		return num1 * num2
	case '/':
		return num1 / num2
	default:
		fmt.Println("运算符错误")
		return -1
	}
}

//返回某个运算符的优先级[程序员定义]
func Priority(oper int) int {
	switch oper {
	case '*', '/':
		return 2
	case '+':
		return 0
	case '-':
		return 1
	default:
		fmt.Println("计算符有问题")
		return -1
	}
}

func main() {
	numStack := &StackStruct{
		Cap: 20,
		Top: -1,
	}
	numStack.Arr = make([]int, numStack.Cap)
	operStack := &StackStruct{
		Cap: 20,
		Top: -1,
	}
	operStack.Arr = make([]int, operStack.Cap)
	exp := "1+4*2-1-5/2" //=6
	for i := 0; i < len(exp); i++ {
		ch := int(exp[i]) //字符串
		switch ch {
		case '+', '-', '*', '/':
			if operStack.Top == -1 { //空栈
				operStack.Push(ch)
			} else {
				//如果发现operStack栈顶的运算符的优先级大于等于当前准备入栈的运算符的优先级
				//就从符号栈的pop取出，，并从数栈的pop取出两个数，进行运算，，运算后的结果再重新入栈
				//到数栈，当前符号再入符号栈
				if Priority(operStack.Arr[operStack.Top]) >= Priority(ch) {
					num2, _ := numStack.Pop()
					num1, _ := numStack.Pop()
					oper, _ := operStack.Pop()
					result := Cal(num1, num2, oper)
					//将计算结果重新入数栈
					numStack.Push(result)

				}
				//当前的符号压入符号栈
				operStack.Push(ch)
			}
		default:
			n := 0
			if i > 0 {
				prech := int(exp[i-1])
				switch prech {
				case '+', '-', '*', '/':
					break
				default:
					n, _ = numStack.Pop()
				}
			}
			numStack.Push(n*10 + (ch - 48))
		}
	}
	//如果扫描表达式完毕，依次从符号栈取出符号，然后从数栈取出两个数
	//运算后的结果，入数栈，直到符号栈为空
	for {
		if operStack.Top == -1 {
			break
		}
		num2, _ := numStack.Pop()
		num1, _ := numStack.Pop()
		oper, _ := operStack.Pop()
		result := Cal(num1, num2, oper)
		//将计算结果重新入数栈
		numStack.Push(result)
	}
	//如果算法没有问题，表达式也是正确的，则结果就是numStack最后数
	// res, _ := numStack.Pop()
	fmt.Printf("表达式%v=%v\n", exp, numStack.Arr[numStack.Top])
}
