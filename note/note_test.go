package note

import "testing"

//测试用函数结构必须叫func Test被测试函数名(log *testing.T) {}
func TestTestableFunction(log *testing.T) {
	testRes := TestableFunction(10)
	if testRes != 12 {
		log.Fatalf("执行错误")
	} else {
		log.Logf("正常执行")
	}

}

//调用方法
//1.go test 正确无日志，错误则输出日志
//2.go test -v 无论正确还是错误都有日志
//3.go test -v note_test.go note.go 只测试这个指定文件
//4.go test -v -test.run Test被测试函数名  只测试这个指定函数
