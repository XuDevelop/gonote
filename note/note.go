package note

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

//文件管理初始化（在项目的根目录下）：go mod init 项目名

//1.1 转义符
func EscapeCharacters() {
	fmt.Print(
		`\n=newline 换行`, "\n", "\n",
		`\r=rarriage return 回车`, "bbbbbbbb\raaa", "\n",
		`\t=tab 制表符`, "\t", "\n",
		`\a=alert sound 警告声`, "a", "\n",
		`\b=back 退格`, "b", "\n",
		`\\=\ backslash 斜线`, "\\", "\n",
		`块注释：
		/*
		fmt.Println("...")
		fmt.Println("...")
		fmt.Println("...")
		*/
		`, "\n",
	)
}

//1.2 变量常量
func VariableAndConstant() {
	fmt.Println("常量变量")
	const (
		a    = 20
		b                 //跟a的值一样
		c    = iota       //行数是从0开始的，所以c在第2行，c=2
		d                 //与c的定义一样，全写是d=iota，所以d=3
		_                 //省略这一行
		f                 //依照行数，f=5
		g, h = iota, iota //并行所以一起等于6
	)
	fmt.Println(a, b, c, d, f, g, h)
}

//2 格式化输出 fmt.Printing Verb
func FmtPrintingVerb() {
	fmt.Println("//2. fmt.Printing Verb")
	n := 10
	var m bool
	//2.1通用
	fmt.Println("//2.1 General")
	fmt.Printf("%%v=value,值的默认格式,%v\n", n)
	fmt.Printf("%%T=Type,%T\n", n)
	fmt.Printf("%%%%=百分号\n")
	//2.2布尔值
	fmt.Println("//2.2 Boolean")
	fmt.Printf("%%t=布尔值，true or false,%t\n", m)
	//2.3整数
	fmt.Println("//2.3 integer")
	fmt.Printf("%%b=binary 二进制,%b\n", n)
	fmt.Printf("%%c=the character by Unicode 传出该unicode值对应的字符 ,%c\n", n)
	fmt.Printf("%%d=decimalism 十进制,%d\n", n)
	fmt.Printf("%%o=octonary number 八进制,%o\n", n)
	fmt.Printf("%%q=with single-quoted 将该值用单引号括起来直接输出,%q\n", n)
	fmt.Printf("%%x=hexadecimal with a-f 十六进制 使用a-f,%x\n", n)
	fmt.Printf("%%X=hexadecimal with A-F 十六进制 使用A-F,%X\n", n)
	fmt.Printf("%%U=Unicode码值 ,%U\n", n)
	//2.4浮点数与复数
	fmt.Println("//2.4 Floating point numbers and complex numbers")
	var a float32 = 11.111111111111111
	fmt.Printf("%%b=无小数部分、二进制指数的科学计数法 ...p-..,%b\n", a)
	fmt.Printf("%%e=科学计数法 e+ （十进制的）,%e\n", a)
	fmt.Printf("%%E=科学计数法 E+ （十进制的）,%E\n", a)
	fmt.Printf("%%x=十六进制的科学计数法 a-f 格式为：0x1.    p+   ,%x\n", a)
	fmt.Printf("%%X=十六进制的科学计数法 A-F 格式为：0x1.    p+   ,%X\n", a)
	fmt.Printf("%%f=float point  输出浮点数的值（无默认宽度，默认精度6）,%f\n", a)
	fmt.Printf("%%F=等价于F,%F\n", a)
	fmt.Printf("%%g=根据实际情况采用%%e或者%%f格式,%g\n", a)
	fmt.Printf("%%G=根据实际情况采用%%E或者%%F格式,%G\n", a)
	fmt.Printf("%%9f=宽度9 默认精度 ,%9f\n", a)
	fmt.Printf("%%.2f=精度2 ,%.2f\n", a)
	fmt.Printf("%%9.2f=宽度9 精度2 ,%9.2f\n", a)
	fmt.Printf("%%9.f=宽度9 精度0 ,%9.f\n", a)
	//2.5字符串和[]byte
	fmt.Println("//2.5 String and []byte")
	var str string = "hello"
	fmt.Printf("%%s=string,%s\n", str)
	fmt.Printf("%%q=with quoted 该值用双引号括起来直接输出,%q\n", str)
	fmt.Printf("%%x=hexadecimal ,lower-case ,two characters per byte 每个字节用两个十六进制字符表示（小写的）,%x\n", str)
	fmt.Printf("%%X=hexadecimal ,lower-case ,two characters per byte 每个字节用两个十六进制字符表示（大写的）,%X\n", str)
	//2.6指针
	fmt.Println("//2.6 Pointer")
	var p *int
	fmt.Printf("%%p=pointer 指针 开头是0x ,%p\n", p)
}

//3 流程控制
//3.2 随机数生成
func RandomNumber() {
	fmt.Println("1.3 随机数生成")
	var randseed = time.Now().UnixNano()
	for i := 1; i < 10; i++ {
		randseed++
		rand.Seed(randseed)
		randNum := rand.Int31n(100)
		fmt.Println(randNum)
	}
}

//3.3 label, 配合goto跳转，break，和continue
func LabelTest() {
	fmt.Println("1")
	fmt.Println("2")
	var n = 0
	if n == 0 {
		goto labelName1
	}
	fmt.Println("3")
	fmt.Println("4")
labelName1:
	fmt.Println("5")
	fmt.Println("6")
labelName2:
	for i := 7; ; {
	labelName3:
		for ; i < 120; i++ {

			if i == 10 {
				continue labelName3
			}
			fmt.Println(i)
			if i == 20 {
				break labelName2
			}
		}
	}
}

//4 匿名函数
func AnonymousFunction() {
	fmt.Println("4 匿名函数")
	fang := func(n1 int, n2 int) int {
		return n1 + n2
	}(1, 2)
	xj := fang
	fmt.Println(xj)
	fmt.Println(fang) //两者输出结果一样的

	//闭包
	fmt.Println("闭包：一个函数与其相关的引用环境组合的一个整体（实体）")
	fang2 := func() func() int {
		n3 := 1
		fmt.Println("hello")
		return func() int {
			n3++
			return n3
		}
	}
	xj2 := fang2()
	//fmt.Println(fang2)
	//fmt.Println(xj2)
	fmt.Printf("%T", xj2)
	fmt.Println(xj2())

}

//5 基本数据类型
//5.1 基本数据类型和string的转换
func BasicDataTypeAndStringConversion() {
	//5.1.1 基本数据类型转成string
	//方法一：fmt.Sprintf("%参数",表达方式) Sprintf根据format参数生成格式化的字符串并返回该字符串
	//func SPrintf(format string, a...interface{}) string
	fmt.Println(`方法一：fmt.Sprintf("%参数",表达方式)`)
	var num1 int = 99
	var num2 float32 = 11.234444
	var num3 uint = 12126127612762
	var b bool = true
	var a byte = 'h'
	var str string
	str = fmt.Sprintf("%d", num1)
	fmt.Println(`str=fmt.Sprintf("`, `%d`, `",num1)`, "输出的效果：", str)
	str = fmt.Sprintf("%f", num2)
	fmt.Println(`"str=fmt.Sprintf("`, `%f`, `",num2)"`, "输出的效果：", str)
	str = fmt.Sprintf("%t", b)
	fmt.Println(`str=fmt.Sprintf("`, `%t`, `",b)`, "输出的效果：", str)
	str = fmt.Sprintf("%c", a)
	fmt.Println(`str=fmt.Sprintf("`, `%c`, `",a)`, "输出的效果：", str)
	//方法二：使用strconv包的函数
	fmt.Println("方法二：使用strconv包的函数")
	//func FormatBool(b bool) string
	str = strconv.FormatBool(b)
	fmt.Printf("var str string ; str=strconv.FormatBool(b) ; 最终输出：str=%s\n", str)
	//func FormatFloat(f float64,fmt byte,prec,bitSize int)string
	//fmt表示格式：'f','b','e','g','E','G'
	//prec 控制精度，对feE表示小数点后的数字个数，对gG控制总的数字个数。
	//如果prec是-1，则代表使用最少数量的，但有必须的数字来表示用f。64:表示这个小数是float64
	str = strconv.FormatFloat(float64(num2), 'f', 4, 32)
	fmt.Printf("str=strconv.FormatFloat(float64(num2),'f',10,10) ; 输出的效果：%s\n", str)
	//func Itoa(i int)string 直接把一个int转成字符串
	str = strconv.Itoa(num1)
	fmt.Printf("str=strconv.Ttoa(num1),输出的效果：%s\n", str)
	//func FormatInt(int64,base int)string     base int 指转成多少进制的值(2~36)
	str = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str=strconv.FormatInt(int64(num1),10) 输出的效果：%s\n", str)
	//func FormatUint(i uint64,base int)string
	str = strconv.FormatUint(uint64(num3), 10)
	fmt.Printf("str=strconv.FormatUint(uint64(num3),10)输出的效果：%s\n", str)

	//5.1.2 string转成基本数据类型
	//方法一：使用strconv包的函数
	//func ParseBool(str string)(value bool,err error)
	str = "true"
	bool1, _ := strconv.ParseBool(str)
	fmt.Println("str=\"true\";bool1,_:=strconv.ParseBool(str) ;输出:\n", bool1)
	//func ParseFloat(s string,bitSize int)(f float64,err error)
	str = "12.33324324124132"
	float1, _ := strconv.ParseFloat(str, 32)
	fmt.Println("str=\"12.33324324124132\"; float1,_:=strconv.ParseFloat(str,16);输出：\n", float1)
	//func ParseInt(s string,base int,bitSize int)(i int64,err error)
	//base int:转成几进制    bitSize int:转成多少，有0,8,16,32,64
	str = "123213213232434"
	int1, _ := strconv.ParseInt(str, 16, 16)
	fmt.Println("str=\"123213213232434\";int1,_:=strconv.ParseInt(str,16,16);输出：\n", int1)
	//func ParseUint(s string,base int,bitSize int)(n uint64,err error)
	str = "21324324"
	uint1, _ := strconv.ParseUint(str, 2, 2)
	fmt.Println("str=\"21324324\";uint1,_:=strconv.ParseUint(str,2,2);输出：\n", uint1)
}

//5.2 指针pointer
func Pointer() {
	fmt.Println("4 pointer")
	//基本数据类型，变量存的就是值，也叫值类型
	//获取变量的地址，用& ，如：var num int,获取num 的地址：&num
	//指针类型，变量存的是一个地址，这个地址指向的空间才是值，如：var ptr *int=&num
	//获取指针类型所指向的值，使用：* ，如：var ptr *int,使用*ptr获取p指向的值
	//下面的var ptr *int=&i
	var i int = 10
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr的地址=%v\n", &ptr)
	fmt.Printf("ptr指向的值=%v\n", *ptr)
	//
	var num *int = &i
	*num = 1
	fmt.Println("i=", i)
}

//5.3 值类型和引用类型

//值类型：基本数据类型int系列，float系列，bool，string、数组和结构体
//变量直接存储值，内存通常在栈中分配
//（内存分栈区和堆区。）
//栈区：值类型数据,通常在栈区
//堆区：引用类型，通常在堆区分配空间
//引用类型：指针、slice切片、map、管道chan、interface。。。
//变量存储是一个地址，这个地址对应的空间才是真正存储数据（值），内存通常在堆区上分配，当没有任何变量引用这个地址时，
//改地址对应的数据空间就成为了一个垃圾，由GC来回收

//6 系统内建的常用函数
//6.1 字符串的常用函数
func FunctionsInStrings() {
	fmt.Println("6.1.1 统计字符串的长度，按字节 len(str) : func len(v Type) int ")
	//数组：v中元素的数量  数组指针：*v中元素的数量（v为nil时panic）
	//切片、映射：v中元素的数量；若v为nil，len（v）即为零
	//字符串：v中字节的数量    通道：通道缓存中队列（未读取）元素的数量；若v为nil，len（v）即为零
	//var a int =123456
	var b string = "hello"
	//a1:=len(a)
	b1 := len(b)
	//fmt.Println("var a int =123456 ; a1:=len(a) ; 输出：",a1)
	fmt.Println("var b string =\"hello\" ; b1:=len(b) ; 输出：", b1)

	fmt.Println("6.1.2 字符串遍历，同时处理有中文的问题 r：=[]rune(str)")
	str2 := "hello你"
	for i := 0; i < len(str2); i++ {
		fmt.Println("字符=", str2[i])
	} //打印的全是数字
	str3 := []rune(str2) //字符串遍历，同时处理有中文的问题 r：=[]rune(str)
	for i := 0; i < len(str3); i++ {
		fmt.Printf("字符=%c\n", str2[i])
	} //打印出字符
	fmt.Println(str3)

	fmt.Println("6.1.3 字符串转整数：func Atoi(s string)(i int,err error)")
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换失败", err)
	} else {
		fmt.Println("转换成功")
		fmt.Printf("Type:%T,value:%v\n", n, n)
	}

	fmt.Println("6.1.4 整数转字符串：str=strconv.Itoa(123)")
	str4 := strconv.Itoa(123)
	fmt.Printf("Type:%T,value:%v\n", str4, str4)

	fmt.Println("6.1.5 字符串转[]byte：var byte=[]byte(\"hello word\")")
	var bytes = []byte("hello word")
	fmt.Printf("Type=%T,char=%c\n", bytes, bytes)

	fmt.Println("6.1.6 []byte转字符串: str=string([]byte{11,22,33})")
	str6 := string([]byte{11, 22, 33})
	fmt.Printf("Type=%T,int=%v\n", str6, str6)

	fmt.Println("6.1.7 查找字符串是否在指定的字符串中：strings.Contains(\"baby\",\"honey\")")
	//func Contains(s,suber string)bool  如果没有指定的字符则输出false
	c := strings.Contains("abc", "ad")
	fmt.Println(c)

	fmt.Println("6.1.8 统计一个字符里有几个指定的字符串：strings.Count(\"abcc\",\"c\")") //区分大小写的
	d := strings.Count("abssBss", "b")
	fmt.Println(d)

	fmt.Println("6.1.9 区分和不区分大小写的字符串比较")
	fmt.Println(strings.EqualFold("abc", "ABC")) //不区分大小写
	fmt.Println("abc" == "ABC")                  //区分大小写

	fmt.Println("6.1.10 返回子串在字符第一次出现的index值，如果没有返回-1  strings.Index(\"NMnhh abc\",\"abc\")")
	//strings.Index("NMnhh abc","abc")
	fmt.Println(strings.Index("NMnhh_abc", "abc")) //6

	fmt.Println("6.1.11 返回子串在字符最后一次出现的index值，如果没有返回-1   strings.LastIndex(\"NMnhh abc\",\"abc\")")
	fmt.Println(strings.LastIndex("NMnhh_abcabcabcabcabc", "abc")) //18

	fmt.Println("6.1.12 将指定的子串替换成另一个子串：strings.Replace(\"go go hello\",\"go\",\"go语言\",n)")
	// fmt.Println(strings.Replace("go go hello","go","go语言",n))   n 可以指定你希望替换几个 n=-1是全部替换
	fmt.Println(strings.Replace("go go hello", "go", "go语言", -1))

	fmt.Println("6.1.13 按照指定的某个字符为分割标识，将字符串拆分成字符串数组  strings.Split(\"hello word,ok\",\",\")")
	zf := strings.Split("hello word,ok", ",")
	for i := 0; i < len(zf); i++ {
		fmt.Println(zf[i])
	}
	fmt.Println(strings.Split("hello word,ok", ","))

	fmt.Println("6.1.14 将字符串的字母进行大小写的转换  strings.ToLower(\"Go\")")
	//strings.ToLower("Go")//go
	//strings.ToUpper("Go")//GO
	fmt.Println(strings.ToLower("Go"))
	fmt.Println(strings.ToUpper("Go"))

	fmt.Println("6.1.15 将字符串左右两边的空格去掉：")
	q := strings.TrimSpace("  aa ss   aa aa   a a a  ")
	fmt.Println(q)

	fmt.Println("6.1.16 将字符串左右两边指定的字符去掉：")
	w := strings.Trim("!aa ss!!aa aa!   a !!a a !!", "!")
	fmt.Println(w)

	fmt.Println("6.1.17 将字符串左边指定的字符去掉：")
	e := strings.TrimLeft("!aa ss!!aa aa!   a !!a a !! ", "!")
	fmt.Println(e)

	fmt.Println("6.1.18 将字符串右边指定的字符去掉：")
	r := strings.TrimRight("!aa ss!!aa aa!   a !!a a !!", "!")
	fmt.Println(r)

	fmt.Println("6.1.19 判断字符串是否以指定的字符串开头")
	t := strings.HasPrefix("abc,wwes22", "abc")
	fmt.Println(t)

	fmt.Println("6.1.20 判断字符串是否以指定的字符串结束")
	y := strings.HasSuffix("sdwdewdxwccmm", "m")
	fmt.Println(y)
}

//6.2 时间和日期函数
func TimeAndData() {
	fmt.Println("6.2 时间和日期函数")
	//6.2.1完整版
	fmt.Println("当前年：", time.Now().Year())
	fmt.Println("当前月：", time.Now().Month())
	fmt.Println("当前月：", int(time.Now().Month())) //可将返回的mouth转成对应的数值
	fmt.Println("当前日：", time.Now().Day())
	fmt.Println("当前日：", time.Now().Hour())
	fmt.Println("当前日：", time.Now().Minute())
	fmt.Println("当前日：", time.Now().Second())

	//格式化当前日期时间
	//方法1
	now := time.Now()
	fmt.Printf("Type=%T,value=%v\n", now, now)
	fmt.Printf("当前年月日：%d-%d-%d %d-%d-%d\n", now.Year(), now.Month(),
		now.Day(), now.Hour(), now.Minute(), now.Second())
	//方法2
	fmt.Println(time.Now().Format("2006/01/02 15:04:05\n")) //是golang的时间点
	fmt.Println(time.Now().Format("2006-01-02\n"))
	fmt.Println(time.Now().Format("15:04:05\n"))
	fmt.Println(time.Now().Format("01\n")) //只显示月份

	//6.2.2时间的常量
	// const(
	// 	Nanosecond Duration=1//纳秒        //Duration是int64的别名
	// 	Microsecond =1000*Nanosecond//微秒
	// 	Millisecond =1000*Microsecond//毫秒
	// 	Second=1000*Millisecond//秒
	// 	Minute=60*Second//分钟
	// 	Hour=60*Minute//小时
	// )

	//6.2.3 休眠
	//func Sleep(d Duration)
	//time.Seelp(100*time.Millisecond)//休眠100毫秒
	//案例：
	for i := 1; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(5 * time.Second)
	}

	//6.2.4 获取当前unix时间戳和unixnano时间戳   //UTC 1970-01-01 00:00(开始时间)

}

//6.3 new(Type)
func FunctionNew() {
	fmt.Println("6.3 new(Type)")
	a := new(int) //传出地址跟var a *int差不多
	fmt.Printf("Type=%T,valeu=%v", a, a)
}

//6.4 错误捕捉延迟处理  //出错后，会输出错误但是整个程序继续执行
func SearchDaly() {
	fmt.Println("6.4 错误捕捉延迟处理  //出错后，会输出错误但是整个程序继续执行")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("成功")
		}
	}()
	i := 1
	a := 0
	fmt.Println(i / a)
	fmt.Println("函数再继续执行")
}

//6.5 自定义错误
func CustomError() {
	fmt.Println("6.5 自定义错误")
	err := func(str string) (err error) {
		if str == "xujing" {
			return nil
		} else {
			return errors.New("你是猪")
		}
	}("fangfang")
	if err != nil {
		panic(err)
	}
}

//7  数组
func Array() {
	//7.1 申明数组
	//第一种：
	var arr [2]int
	arr[0] = 0
	arr[1] = 1
	fmt.Println(arr)
	//第二种：
	var arr2 = [2]int{2, 3}
	fmt.Println(arr2)
	//第三种：
	var arr3 = [...]int{4, 5}
	fmt.Println(arr3)
	//第四种：
	var arr4 = [...]int{1, 2, 3, 4}
	fmt.Println(arr4) //2:6中，2是下标，意思是第二个数字是6

	//7.2 For...range遍历数组
	for index, value := range arr {
		fmt.Println("第", index, "个数组的值是：", value)
	}

	//7.3 声明切片
	//7.3.1 引用数组
	slice1 := make([]int, 10)
	s := slice1
	slice1[0] = 666
	fmt.Println(s[0])
	//7.3.2 make([]type,len,cap(optional))    len是定义该切片的长度 cap是容量必须>=len
	var slice2 = make([]int, 3) //or slice2:=make([]int,3)
	slice2[0] = 99
	fmt.Println(slice2[0])

	//7.7 append的用法
	//7.4.1 用法一：
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)
	//7.4.2 用法二：
	slice2 = append(slice2, 3)
	fmt.Println(slice2)

	//7.5 copy()各种情况
	copy(slice1, []int{11, 12, 13, 14})
	fmt.Println(slice1)
	var a = []int{1, 2}
	copy(a, []int{9, 9, 9, 9, 9})
	fmt.Println(a)
}

//8 map,key是无序的，map是引用类型
func MapAndKey() {
	//8.1 声明map，make（），和直接声明，赋值
	var map1 = make(map[int]string)
	var map2 = make(map[int]map[int]string)
	map3 := map[int]string{
		1: "Aa",
		2: "Bb",
		3: "Cc",
	}
	map1[1] = "芳芳"
	map1[2] = "他是"
	map1[3] = "lsp"
	map2[4] = make(map[int]string)
	map2[4][1] = "上面是真理！"
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	//8.2 删除元素 delete()
	delete(map1, 2)
	fmt.Println(map1)

	//查找map
	v, ok := map1[1]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("没找到")
	}

	//遍历range
	for k, v := range map3 {
		fmt.Println(k)
		for k2, v2 := range v {
			fmt.Println(k2, ":", v2)
		}
	}
}

//9 创建结构体实例
type student struct {
	Name string
	Age  int
}

func StructureInstance() {
	//方法一：
	var stu1 student = student{"数字", 1}
	stu2 := student{"数字2", 2}
	var stu3 student = student{
		Name: "Anna",
		Age:  20,
	}
	stu4 := student{
		Name: "Anna",
		Age:  20,
	}
	fmt.Println(stu1, stu2, stu3, stu4)
	//方法二(返回结构体的指针类型)：
	var stu5 *student = &student{"aaa", 11}
	var stu6 *student = &student{
		Name: "lisa",
		Age:  18,
	}
	fmt.Printf("%T", stu5)
	fmt.Println(*stu5, stu6) //加个*输出去掉地址符
}

//10 工厂模式编程
//10.1 工厂模式结构体
func NewStudent(n string, a int) *student {
	return &student{
		Name: n,
		Age:  a,
	}
}

//在另外函数里先引包，然后：
// func main(){
// 	var stu = note.NewStudent("tom",18)
// 	fmt.Println(*stu)
//  fmt.Println("name=",stu.Name,"Age=",stu.Age)
// }
type teacher struct {
	Name string
	age  int
}

func NewTeacher(n string, a int) *teacher {
	return &teacher{
		Name: n,
		age:  a,
	}
}
func (t *teacher) GetAge() int {
	return t.age
}

//在另一个函数里调用：
// func main(){
// 	var tea = note.NewTeacher("tom",18)
// 	fmt.Println(*tea)
// 	fmt.Println("name=",tea.Name,"Age=",tea.GetAge())
// }

//11 继承
type Pupil struct {
	student
	Name string
}
type UniversityStudent struct {
	*student
}

func Extend() {
	var student1 Pupil
	student1.Name = "student001"
	student1.student.Name = "student002"
	//var student2 = UniversityStudent{&student{"aaa", 111}}
	var student2 = UniversityStudent{
		&student{
			Name: "aaa",
			Age:  111,
		},
	}
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(*student2.student)
}

//12 接口 interface
type Usb interface {
	//声明两个方法
	Start()
	Stop()
}
type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("phone start......")
}
func (p Phone) Stop() {
	fmt.Println("phone stop......")
}
func (p Phone) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

type Computer struct {
}

func (c *Computer) Start() {
	fmt.Println("computer start......")
}
func (c *Computer) Stop() {
	fmt.Println("computer stop......")
}
func (computer *Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}
func InterfaceTest() {
	//测试
	//先创建结构体变量
	computer := Computer{}
	phone := Phone{}

	computer.Working(phone)
	//芳芳是傻子
}

//13 类型断言
func TypeAssertion() {
	var fang interface{}
	var n1 int = 123
	var n2 interface{} = 99

	v, ok := n2.(int)
	fang = n1 + v
	if !ok {
		fmt.Println("断言出错！")
	} else {
		fmt.Printf("Type of fang =%T,Value of fang =%v\n", fang, fang)
		fmt.Printf("Type of v =%T,Value of v =%v", v, v)
	}
}

//14 文件操作
func FileOperation() {
	filePath := "E:/develop/gonote/file.txt"
	// 14.1 打开与关闭
	file1, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开file.txt文件出错！", err)
		return
	}
	defer func() {
		err = file1.Close()
		if err != nil {
			fmt.Println("关闭file.txt文件出错！", err)
		}
	}()

	//14.2 带缓冲区的读取写文件（大文件）
	//14.2.1 带缓冲区的读取，适合大文件的读取，需要手动打开与关闭
	file2, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//os.OpenFile(文件路径,打开模式（如果多个打开模式要用|分开）,权限代码（只对unix生效）)
	//os.O_RDWR ：read and write
	//os.O_CREATE : 如果文件不存在则创建
	//os.O_APPEND : 写入时追加在结尾
	//os.O_EXCL : excluding 与os.O_CREATE一起使用，文件不能存在
	//os.O_RDONLY : 只能打开读取不能写入
	// os.O_WRONLY :只能写入不能读取
	// os.O_SYNC : open for synchronous I/O  为同步 I/O 打开
	// os.O_TRUNC : 打开时缩短常规可写文件
	//权限代码：
	//7:read and write and run
	//6:read and write
	//5:write and run
	//0:没有任何权限
	//根用户，主用户，访问者
	if err != nil {
		fmt.Println("打开file.txt文件出错！", err)
		return
	}
	defer func() {
		err = file2.Close()
		if err != nil {
			fmt.Println("关闭file.txt文件出错！", err)
		}
	}()

	reader := bufio.NewReader(file2) //创建读文件的缓冲区
	for {
		str, err := reader.ReadString(' ') //文件内容读到括号里的字符就结束读取
		if err != nil {
			if err == io.EOF { // io.EOF : EOF is the error returned by Read when no more input is available 文件已经被读完了
				break
			}
			fmt.Println("读取文件失败", err)
			return
		}
		fmt.Println(str)
	}

	//14.2.2 带缓冲区的写入，适合大文件，需要手动打开与关闭
	str := "123456\n"
	writer := bufio.NewWriter(file2)
	for i := 0; i < 6; i++ {
		_, err := writer.WriteString(str)
		if err != nil {
			fmt.Println("写入缓冲区失败！", err)
			return
		}

	}
	err = writer.Flush() //将缓冲区的数据写入文件
	if err != nil {
		fmt.Println("将缓冲区的数据写入文件失败！", err)
		return
	}

	//14.3.1 一次性读取文件（小文件）
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("一次性读取文件失败！", err)
		return
	}
	fmt.Println("一次性读取到的：", string(content))

	//14.3.2 一次性写入文件（小文件）
	content = []byte("pig" + "fangpig")
	err = ioutil.WriteFile(filePath, content, 0666) //这种写入是直接覆盖整个文件
	if err != nil {
		fmt.Println("一次性写入文件失败！", err)
		return
	}

	//14.4 判断文件是否存在
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("查询文件状态出错！", err)
	} else {
		fmt.Println("文件存在")
	}
	fmt.Println(fileInfo)

	if os.IsNotExist(err) {
		fmt.Println("文件不存在")
	}

	//14.5 拷贝文件
	file3, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开文件失败！", err)
		return
	}
	defer func() {
		err = file3.Close()
		if err != nil {
			fmt.Println("关闭file3.txt文件出错！", err)
		}
	}()

	reader = bufio.NewReader(file3) //创建读文件的缓冲区
	//写一个目标文件路径
	destinationFilePath := "E:/develop/gonote/dst.txt"
	dstFile, err := os.OpenFile(destinationFilePath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		fmt.Println("创建文件失败！", err)
		return
	}
	defer func() {
		err = dstFile.Close()
		if err != nil {
			fmt.Println("关闭dst.txt文件出错！", err)
		}
	}()
	writer = bufio.NewWriter(dstFile)
	written, err := io.Copy(writer, reader)
	if err != nil {
		fmt.Println("拷贝文件失败！", err)
		return
	}
	fmt.Println("拷贝文件成功，拷贝了", written, "个字节")

}

//15 获取和解析命令行参数
//15.1 获取命令行参数
func GetCommandLineArguments() {
	fmt.Printf("接收到命令行参数%v个\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("os.Args[%v]=%v\n", i, v)
	}
}

//15.2 解析命令行参数
func ParseCommandLineArgs() {
	var user string
	var pwd string
	var host string //主机
	var port int    //端口
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	//flag.StringVar(&变量，命令行指定字段，默认值，注释)
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认localhost")
	flag.IntVar(&port, "port", 8080, "端口，默认为8080")
	flag.Parse() //解析
	fmt.Printf("user=%s,pwd=%s,host=%s,port=%v\n", user, pwd, host, port)

	// go run main.go -u sb -port 5656 -h 23.223.10.02 -pwd sdgsjahjsda
}

//16 json的序列化和反序列化
type JsonStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func JsonOperation() {
	var jsonStruct1 JsonStruct = JsonStruct{
		Name: "tom",
		Age:  18,
	}
	// 16.1 序列化
	data, err := json.Marshal(&jsonStruct1)
	if err != nil {
		fmt.Println("json序列化失败！", err)
		return
	}
	fmt.Println("序列化结果为：", string(data))
	//16.2 反序列化
	var jsonStruct2 JsonStruct
	err = json.Unmarshal(data, &jsonStruct2)
	if err != nil {
		fmt.Println("json反序列化失败！", err)
		return
	}
	fmt.Println("反序列化的结果为：", jsonStruct2)
}

//17 单元测试 见note_test.go(文件名必须为：原始文件名_test.go)
//正规工作中所有代码都要进行单元测试
func TestableFunction(i int) int {
	return i + 1
}

//18 协程 goroutine
var (
	goroutineComputRes  []int
	goroutineComputLock sync.Mutex //同步.互斥  可以但不推荐（效率低）
)

func GoroutineFactorialComputer(n int) { //协程积乘计算器
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	goroutineComputLock.Lock()    //锁定
	goroutineComputRes[n-1] = res //存储
	goroutineComputLock.Unlock()  //解锁

}

func RunGoroutineFactorialComputer() {
	//18.1 获取当前计算机逻辑cpu数量
	cpuNum := runtime.NumCPU()
	fmt.Printf("当前计算机有%v个逻辑cpu\n", cpuNum)
	//18.2 自定义系统可以调用的cpu数量，默认调用全部cpu
	if cpuNum > 3 {
		runtime.GOMAXPROCS(cpuNum - 1)
	}
	//18.3 goroutine积乘计算案例
	n := 30
	goroutineComputRes = make([]int, n)
	for i := 1; i <= n; i++ {
		go GoroutineFactorialComputer(i)
	}
	time.Sleep(time.Second * 3)
	fmt.Println(goroutineComputRes)

	//18.4 协程错误处理
	//defer+recover  见434行
}

//19 管道 channel （数据先进先出）
func Channel() {
	//19.1 管道的声明和初始化
	var intChan chan int = make(chan int, 10) //如果不指定容量，默认容量为0
	//19.2 存取数据
	//19.2.1 存入数据
	intChan <- 12
	intChan <- 100
	intChan <- 16
	fmt.Printf("len(intChan)=%v,cap(intChan)=%v\n", len(intChan), cap(intChan))
	//19.2.2 取出数据
	var num int = <-intChan
	<-intChan //不接受，丢弃数据
	fmt.Printf("len(intChan)=%v,cap(intChan)=%v\n", len(intChan), cap(intChan))
	fmt.Println("num=", num)
	//19.3 遍历管道（如果不提前关闭管道，则会报错）
	close(intChan)
	for v := range intChan {
		fmt.Println("v=", v)
	}
	//等效于
	for {
		v, ok := <-intChan //如果数据取完则会返回false
		if !ok {
			break
		}
		fmt.Println("v=", v)
	}
	//19.4 只读只写channel, 结合函数传参使用
	var readOnlyChan chan int
	var writeOnlyChan chan int
	func(r <-chan int, w chan<- int) {
		// <-chan  : 只读
		// chan<- : 只写
	}(readOnlyChan, writeOnlyChan)
	//19.5 selest无法确定何时关闭管道时用
	var strChan1 chan string = make(chan string, 3)
	var strChan2 chan string = make(chan string, 3)
	strChan1 <- "fang"
	strChan2 <- "is"
	strChan1 <- "a"
	strChan2 <- "pig"
	var b bool
	for {
		select {
		case v := <-strChan1:
			fmt.Println(v, " ")
		case v := <-strChan2:
			fmt.Println(v, " ")
		default:
			b = true
		}
		if b {
			break
		}
	}
}

//20 反射
//20.1 反射的常用方法
func CommonMethodsOfReflect(i interface{}) {
	fmt.Println("\n//20.1 反射的常用方法")
	//20.1.1 获取反射类型
	reflectType := reflect.TypeOf(i)
	fmt.Printf("value of reflectType=%v,Type of reflectType=%T\n", reflectType, reflectType)
	//20.1.2 获取反射值
	reflectValue := reflect.ValueOf(i)
	fmt.Printf("value of reflectValue=%v,Type of reflectValue=%T\n", reflectValue, reflectValue)
	//20.1.3 将反射的值转换回空接口
	interfaceValue := reflectValue.Interface()
	fmt.Printf("value of interfaceValue=%v,Type of interfaceValue=%T\n", interfaceValue, interfaceValue)
	//20.1.4 用类型断言 把空接口转回所需类型
	intValue := interfaceValue.(int)
	fmt.Printf("value of intValue=%v,Type of intValue=%T\n", intValue, intValue)

}

//20.2 对指针的反射
func ReflectOnPointer(i interface{}) {
	fmt.Println("\n//20.2 对指针的反射")
	// 获取反射值
	reflectValue := reflect.ValueOf(i)
	fmt.Printf("value of reflectValue=%v,Type of reflectValue=%T\n", reflectValue, reflectValue)
	//20.2.1 修改指针指向的原值
	reflectValue.Elem().SetString("tom")
}

//20.3 对结构体的反射
type StructReflect struct {
	No1 string `json:"no1"`
	No2 int    `json:"no2"`
	No3 bool
	No4 float64
}

func (sr StructReflect) StructMethod1() {
	fmt.Println("StructMethod1()被调用了")
}
func (sr StructReflect) StructMethod2(i int, j int) int {
	fmt.Printf("StructMethod2()被调用了，并接收到了i=%v，j=%v\n", i, j)
	return i + j
}

func ReflectOnStruct(i interface{}) {
	fmt.Println("\n//20.3 对结构体的反射")
	//获取反射类型
	reflectType := reflect.TypeOf(i)
	fmt.Printf("value of reflectType=%v,Type of reflectType=%T\n", reflectType, reflectType)
	//获取反射值
	reflectValue := reflect.ValueOf(i)
	fmt.Printf("value of reflectValue=%v,Type of reflectValue=%T\n", reflectValue, reflectValue)
	//将反射的值转换回空接口
	interfaceValue := reflectValue.Interface()
	fmt.Printf("value of interfaceValue=%v,Type of interfaceValue=%T\n", interfaceValue, interfaceValue)
	//用类型断言 把空接口转回所需类型
	intValue := interfaceValue.(StructReflect)
	fmt.Printf("value of intValue=%v,Type of intValue=%T\n", intValue, intValue)
	//20.3.1 获取reflectValue对应的Kind
	reflectKind := reflectValue.Kind()
	fmt.Printf("value of reflectKind=%v,Type of reflectKind=%T\n", reflectKind, reflectKind)
	if reflectKind != reflect.Struct {
		fmt.Println("调用的ReflectOnStruct()时传入的不是结构体")
		return
	}
	//20.3.2 获取结构体的字段数量
	fieldNum := reflectValue.NumField()
	fmt.Printf("用户传入的结构体有%v个字段\n", fieldNum)
	//20.3.3 遍历结构体字段
	for i := 0; i < fieldNum; i++ {
		fmt.Printf("Field(%v)=\"%v:%v\",", i, reflectType.Field(i).Name, reflectValue.Field(i))
		tagValue := reflectType.Field(i).Tag.Get("json")
		if tagValue != "" {
			fmt.Printf("jsonTag=\"%v\"", tagValue)
		}
		fmt.Println()
	}
	//20.3.4 统计这个结构体的方法数量
	methodNum := reflectValue.NumMethod()
	fmt.Printf("用户传入的结构体有%v个方法\n", methodNum)
	//20.3.5 在反射里调用结构体方法
	reflectValue.Method(0).Call(nil)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(666), reflect.ValueOf(999))
	resOfMethod := reflectValue.Method(1).Call(params)
	for i, v := range resOfMethod {
		fmt.Printf("value of resOfMethod[%v]=%v,Type of resOfMethod[%v]=%T\n", i, v, i, v)
	}
}

//20.4 对结构体指针的反射
func ReflectOnStructPtr(i interface{}) {
	fmt.Println("\n//20.4 对结构体指针的反射")
	//获取反射类型
	reflectType := reflect.TypeOf(i)
	fmt.Printf("value of reflectType=%v,Type of reflectType=%T\n", reflectType, reflectType)
	//20.4.1 获取反射类型指向的类型（原始类型）
	typeByReflectType := reflectType.Elem()
	fmt.Printf("value of typeByReflectType=%v,Type of typeByReflectType=%T\n",
		typeByReflectType, typeByReflectType)
	//获取反射值
	reflectValue := reflect.ValueOf(i)
	fmt.Printf("value of reflectValue=%v,Type of reflectValue=%T\n", reflectValue, reflectValue)
	//20.4.2 创建一个新的typePointedToByReflectType的指针
	newReflectValue := reflect.New(typeByReflectType)
	fmt.Printf("value of newReflectValue=%v,Type of newReflectValue=%T\n", newReflectValue, newReflectValue)
	//20.4.3 获取newReflectValue指向的原始值
	valueByNewReflectValue := newReflectValue.Elem()
	valueByNewReflectValue.FieldByName("No1").SetString("somebody") //根据字段赋值
	valueByNewReflectValue.FieldByName("No2").SetInt(10)
	fmt.Printf("value of valueByNewReflectValue=%v,Type of valueByNewReflectValue=%T\n", valueByNewReflectValue, valueByNewReflectValue)
	//20.4.4 将newReflectValue还原为原始结构体
	newSr := newReflectValue.Interface().(*StructReflect)
	fmt.Printf("value of newSr=%v,Type of newSre=%T\n", newSr, newSr)
}

func Reflect() {
	CommonMethodsOfReflect(10)
	str := "pig"
	strPtr := &str
	fmt.Println("反射修改前的值：", *strPtr)
	ReflectOnPointer(strPtr)
	fmt.Println("通过反射修改后的值：", *strPtr)
	var sr StructReflect = StructReflect{
		No1: "fang",
		No2: 60,
		No3: true,
		No4: 2.333,
	}
	ReflectOnStruct(sr)
	srPtr := &sr
	ReflectOnStructPtr(srPtr)
}

//21 TCP编程 (Transmission Control Protocol 传输控制协议)
func TCPServerProcess(conn net.Conn) {
	//21.1.4 关闭连接
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("关闭链接失败！", err)
		}
	}()
	for {
		fmt.Println("等待客户端发送消息...")
		//21.1.5 读取客户端信息
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端已退出！")
			} else {
				fmt.Println("读取客户端消息失败！", err)
			}
			return
		}
		fmt.Println("读取到客户端的消息：", string(buf[0:n]))
	}
}

//21.1 TCP服务器
func TCPServer() {
	//21.1.1 监听端口
	listener, err := net.Listen("tcp", "0.0.0.0:21104")
	if err != nil {
		fmt.Println("监听失败！", err)
		return
	}
	//21.1.2 关闭监听端口
	defer func() {
		err := listener.Close()
		if err != nil {
			fmt.Println("关闭监听失败！", err)
		}
	}()
	fmt.Println("监听成功！", listener)
	for {
		fmt.Println("等待客户端连接...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("客户端连接失败！", err)
			continue
		}
		fmt.Println("客户端连接成功！", conn)
		fmt.Println("客户端来自：", conn.RemoteAddr().String())
		//21.1.3 启动协程处理
		go TCPServerProcess(conn)
	}

}

//21.2 TCP客户端
func TCPClient() {
	//21.2.1 拨号连接
	conn, err := net.Dial("tcp", "localhost.betadevelop.com:21104")
	if err != nil {
		fmt.Println("拨号连接失败！", err)
		return
	}
	//21.2.2 关闭连接
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("关闭连接失败！", err)
		}
	}()
	fmt.Println("拨号连接成功！", conn)
	for {
		fmt.Println("请输入要发送的信息：")
		readerPtr := bufio.NewReader(os.Stdin)
		str, err := readerPtr.ReadString('\n')
		if err != nil {
			fmt.Println("读取失败！", err)
			continue
		}
		str = strings.Trim(str, "\r\n")
		if str == "exit" {
			return
		}
		//21.2.3 发送数据
		n, err := conn.Write([]byte(str))
		if err != nil {
			fmt.Println("发送失败！", err)
			return
		}
		fmt.Println("发送成功，发送了", n, "个字节")
	}
}

//22 redis数据库
//redis支持的数据类型: String/Hash(类似于golang的map[string]string)/List/Set(无序且不能重复)/zset是有序的(无序：set)

//22.1 常见key value操作
//22.1.1 添加和修改key value：set key value
//22.1.2 添加临时的key value：setex key seconds value
//22.1.3 批量添加修改key value：mset key1 value1 key2 value2 ...
//22.1.4 查看所有的key：keys *
//22.1.5 查看key对应的值：get key
//22.1.6 批量获取key的值：mget key1 key2 ...
//22.1.7 切换数据库（0~15 默认为0）：select index
//22.1.8 查看当前数据库key value数量：dbsize
//22.1.9 清空当前数据库所有key value：flushdb
//22.1.10 清空所有数据库所有key value：flushall
//22.1.11 删除指定的key value：del key

//22.2 hash 操作
//22.2.1 添加修改hash：hset key field value
//22.2.2 批量添加修改hash：hmset key field1 value1 field2 value2 ...
//22.2.3 查看key field对应的值：hget key field
//22.2.4 批量查看key的多个field的值：hmget key field1 field2 ...
//22.2.5 获取key对应的所有的field value: hgetall key
//22.2.6 删除对应的field和value：hdel key field1 field2 ...
//22.2.7 查看hash长度：hlen key
//22.2.8 查看字段是否存在：hexists key field

//22.3 list
//22.3.1 从左边插入元素：lpush key value2 value1 value0 ...
//22.3.2 从右边插入元素：rpush key value0 value1 value2 ...
//22.3.3 获取对应的元素：lrange key startIndex stopIndex (index可以为负数：-1为倒数第一个元素，-2为倒数第二个元素...)
//22.3.4 从左边推出元素（如果全部推出，则对应的key也会被删除）：lpop key
//22.3.5 从右边推出元素：rpop key
//22.3.6 统计list长度：llen key

//22.4 set
//22.4.1 添加元素：sadd key member1 member2 member0 ...
//22.4.2 获取所有的元素：smembers key
//22.4.3 判断元素是否存在：sismember key member
//22.4.4 删除元素：srem key member member ...

//22.5 通过redigo 操作redis
func RediGo() {
	//22.5.1 拨号连接
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis连接失败！", err)
		return
	}
	//22.5.2 关闭连接
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("关闭与redis的连接失败！", err)
		}
	}()
	fmt.Println("redis连接成功！", conn)

	//22.5.3 操作redis
	conn.Do("set", "name", "fangpig")
	str, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("操作redis失败！", err)
		return
	}
	fmt.Println("redis返回：", str)
	conn.Do("hmset", "user1", "name", "fangpig", "age", "60")
	strs, err := redis.Strings(conn.Do("hgetall", "user1"))
	if err != nil {
		fmt.Println("操作redis失败！", err)
		return
	}
	for i, v := range strs {
		if i%2 == 0 {
			fmt.Printf("user1[\"%v\"]:", v)
		} else {
			fmt.Printf("\"%v\"\n", v)
		}
	}

}

//22.5.4 redis连接池（服务器端并发性能优化）
var redisPool *redis.Pool

func RedisPoolInit() { //需要在项目的init(初始化的意思）中调用
	redisPool = &redis.Pool{
		MaxIdle:     8,  //最大空闲连接数
		MaxActive:   0,  //最大连接数，0是没有限制
		IdleTimeout: 30, //最大空闲时间（秒）
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}
func RedisPoolTest() { //通常在协程中使用
	//从redis中取出连接
	conn := redisPool.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("关闭与redis的连接失败！", err)
		}
	}()
	fmt.Println("redis连接成功！", conn)
	conn.Do("hmset", "user1", "name", "fangpig", "age", "60")
	strs, err := redis.Strings(conn.Do("hgetall", "user1"))
	if err != nil {
		fmt.Println("操作redis失败！", err)
		return
	}
	for i, v := range strs {
		if i%2 == 0 {
			fmt.Printf("user1[\"%v\"]:", v)
		} else {
			fmt.Printf("\"%v\"\n", v)
		}
	}
}

//23 棋盘
type ArrayStr struct {
	R int //row 行
	C int //col 列
	V int
}
type PutArrayStr struct {
	SparseArray []ArrayStr
	Len         int
}

func ChessBoard() {
	//1.先创建一个原始数组
	var array [13][10]int
	array[1][2] = 1
	array[2][3] = 2
	//2.输出看看原始的数组
	fmt.Println("原始数组：")
	for _, v := range array {
		for _, v2 := range v {
			fmt.Print("\t", v2)
		}
		fmt.Println()
	}
	//3.转成稀疏数组
	var sparseArray []ArrayStr
	//标准的一个稀疏数组应该还有一个 记录元素的二维数组的规模（行和列，默认值）
	arrayStr := ArrayStr{
		R: len(array),
		C: len(array[0]),
		V: 0,
	}
	sparseArray = append(sparseArray, arrayStr)
	for i, v := range array {
		for i2, v2 := range v {
			if v2 != 0 {
				//创建一个ValNode值结点
				arrayStr := ArrayStr{
					R: i,
					C: i2,
					V: v2,
				}
				sparseArray = append(sparseArray, arrayStr)
			}
		}
	}
	//输出稀疏数组
	fmt.Println("稀疏数组：") //sparseArray 稀疏数组
	for i := range sparseArray {
		fmt.Printf("sparseArray[%d]=r:%d\tc=%d\tv=%d\n", i, sparseArray[i].R, sparseArray[i].C, sparseArray[i].V)
	}

	//把稀疏数组存入硬盘
	//从硬盘读取并还原
	putArrayStr := PutArrayStr{
		SparseArray: sparseArray,
		Len:         len(sparseArray),
	}
	content, _ := json.Marshal(putArrayStr)
	err := ioutil.WriteFile("E:/develop/gonote/note/note.no", content, 0666) //这种写入是直接覆盖整个文件
	if err != nil {
		fmt.Println("一次性写入文件失败！", err)
		return
	}
	content, err = ioutil.ReadFile("E:/develop/gonote/note/note.no")
	if err != nil {
		fmt.Println("一次性读取文件失败！", err)
		return
	}
	fmt.Println("一次性读取到的：", string(content))
	putArrayStr2 := PutArrayStr{}
	json.Unmarshal(content, &putArrayStr2)

	var RArray [][]int = make([][]int, putArrayStr2.SparseArray[0].R)
	for i := range RArray {
		RArray[i] = make([]int, putArrayStr2.SparseArray[0].C)
	}
	for i, v := range putArrayStr2.SparseArray {
		if i == 0 {
			continue
		}
		RArray[v.R][v.C] = v.V
	}

	fmt.Println("还原以后的二维切片:")
	for _, v := range RArray {
		for _, v2 := range v {
			fmt.Printf("%v\t", v2)
		}
		fmt.Println()
	}

}

//24 数组虚拟队列
type QueueStr struct {
	MaxSize int
	Array   [5]int //数组=》模拟队列
	Front   int    //表示指向队列首
	Rear    int    //表示指向队列尾
}

//添加数据到队列
func (q *QueueStr) AddQueue(num int) (err error) {
	//先判断队列是否已满
	if q.Rear == q.MaxSize-1 {
		//Rear是队列尾部（含最后的元素）
		return errors.New("Queue full")
	}
	q.Rear++ //Rear后移
	q.Array[q.Rear] = num
	return
}

func (q *QueueStr) GetQueue() (num int, err error) {
	//先判断队列是否为空
	if q.Rear == q.Front {
		//队列为空
		return -1, errors.New("Queue Empyt")
	}
	q.Front++
	num = q.Array[q.Front]
	return num, err
}

//显示队列，找到队首，然后遍历到队尾
func (q *QueueStr) ShowQueue() {
	fmt.Println("队列当前的情况是：")
	//q.Front不包含队首的元素
	for i := q.Front + 1; i <= q.Rear; i++ {
		fmt.Printf("Array[%d]=%d\t", i, q.Array[i])
	}
	fmt.Println()
}
func Queue() {
	//创建一个队列
	queue := &QueueStr{
		MaxSize: 5,
		Front:   -1,
		Rear:    -1,
	}
	var key int
	var num int
	for {
		fmt.Println("1-加入数据到队列")
		fmt.Println("2-从队列取出数据")
		fmt.Println("3-显示队列")
		fmt.Println("4-退出")
		fmt.Println("输入（1~3）：")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("请输入要添加的一个数据：")
			fmt.Scanln(&num)
			err := queue.AddQueue(num)
			if err != nil {
				fmt.Println("添加失败！ err=", err.Error())
			} else {
				fmt.Println("添加成功！")
			}
		case 2:
			num, err := queue.GetQueue()
			if err != nil {
				fmt.Println("取出失败！ err=", err.Error())
			} else {
				fmt.Println("取出成功，num=", num)
			}
		case 3:
			queue.ShowQueue()
		case 4:
			os.Exit(0)
		default:
			return
		}
	}
}

//25 闭环队列
type CircleQueue struct {
	MaxSize int //4
	Arrary  [4]int
	Head    int
	Tail    int
}

//放入
func (c *CircleQueue) Push(val int) (err error) {
	if c.IsFull() {
		return errors.New("queue full")
	}
	c.Arrary[c.Tail] = val
	c.Tail = (c.Tail + 1) % (c.MaxSize + 1)
	return nil
}

//取出
func (c *CircleQueue) Pop() (val int, err error) {
	if c.IsEmpty() {
		return -1, errors.New("queue empty")
	}
	val = c.Arrary[c.Head]
	c.Head = (c.Head + 1) % (c.MaxSize + 1)
	return
}

//显示
func (c *CircleQueue) Show() {
	size := c.Size()
	if size == 0 {
		fmt.Println("队列为空")
		return
	}
	t := c.Head
	for i := 0; i < size; i++ {
		fmt.Printf("Arrary[%d]=%d\n", t, c.Arrary[t])
		t = (t + 1) % c.MaxSize
	}
}

//判断环形队列是否为满
func (c *CircleQueue) IsFull() bool {
	return (c.Tail+1)%c.MaxSize == c.Head
}

//判断环形队列是否为空
func (c *CircleQueue) IsEmpty() bool {
	return c.Tail == c.Head
}

//判断环形队列有多少个元素
func (c *CircleQueue) Size() int {
	return (c.Tail + c.MaxSize - c.Head) % c.MaxSize
}

func CircleQueueTest() {
	queue := &CircleQueue{
		MaxSize: 5,
	}
	var key int
	var num int
	for {
		fmt.Println("1-加入数据到队列")
		fmt.Println("2-从队列取出数据")
		fmt.Println("3-显示队列")
		fmt.Println("4-退出")
		fmt.Println("输入（1~4）：")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("请输入要添加的一个数据：")
			fmt.Scanln(&num)
			err := queue.Push(num)
			if err != nil {
				fmt.Println("添加失败！ err=", err.Error())
			} else {
				fmt.Println("添加成功！")
			}
		case 2:
			num, err := queue.Pop()
			if err != nil {
				fmt.Println("取出失败！ err=", err.Error())
			} else {
				fmt.Println("取出成功，num=", num)
			}
		case 3:
			queue.Show()
		case 4:
			os.Exit(0)
		default:
			return
		}
	}
}

//26 链表
//26.1.1 单链表
type UnidirectionalChainNode struct {
	No   int
	Name string
	Next *UnidirectionalChainNode //表示指向下一个结点
}

func (head *UnidirectionalChainNode) Append(newScn *UnidirectionalChainNode) {
	t := head
	for {
		if t.Next == nil {
			break
		}
		t = t.Next
	}
	t.Next = newScn
}

//显示链表的所有结点信息
func (head *UnidirectionalChainNode) Print() {
	t := head
	if t.Next == nil {
		fmt.Println("该链表为空")
		return
	}
	for {
		fmt.Printf("UnidirectionalChainNode[%d]=Name:%s==>\t", t.Next.No, t.Next.Name)
		t = t.Next
		if t.Next == nil {
			fmt.Println()
			break
		}
	}
}

//26.1.2 按序号插入节点
func (head *UnidirectionalChainNode) Insert(newScn *UnidirectionalChainNode) {
	t := head
	for {
		if t.Next == nil {
			break
		}
		if t.Next.No > newScn.No {
			break
		} else if t.Next.No == newScn.No {
			fmt.Println("已经存在这个No了：", t.Next)
			return
		}
		t = t.Next
	}
	newScn.Next = t.Next
	t.Next = newScn
}

//26.1.3 按序号删除节点
func (head *UnidirectionalChainNode) Del(no int) {
	t := head
	for {
		if t.Next == nil {
			fmt.Println("不存在这个No")
			return
		}
		if t.Next.No == no {
			break
		}
		t = t.Next
	}
	t.Next = t.Next.Next
}

//测试UnidirectionalChainNode
func UnidirectionalChainNodeTest() {
	//1.创建一个头结点
	head := &UnidirectionalChainNode{}

	//2.创建一个新的UnidirectionalChainNode
	sc1 := &UnidirectionalChainNode{
		No:   1,
		Name: "小明",
	}
	sc2 := &UnidirectionalChainNode{
		No:   2,
		Name: "小红",
	}
	sc3 := &UnidirectionalChainNode{
		No:   3,
		Name: "小兰",
	}
	head.Append(sc1)
	head.Append(sc2)
	head.Append(sc3)
	head.Print()
	head.Del(0)
	head.Print()
}

//26.2.1 新增双向链表节点的结构体
type BidirectionalChainNode struct {
	No   int
	Data string
	Pre  *BidirectionalChainNode //表示指向前一个结点
	Next *BidirectionalChainNode //表示指向后一个结点
}

//26.2.4
func (head *BidirectionalChainNode) Append(newBCN *BidirectionalChainNode) {
	t := head
	for {
		if t.Next == nil {
			break
		}
		t = t.Next
	}
	t.Next = newBCN
	newBCN.Pre = t
}

//26.2.5
func (head *BidirectionalChainNode) Print() {
	t := head
	if t.Next == nil {
		fmt.Println("该链表为空")
		return
	}
	for {
		fmt.Printf("BidirectionalChainNode[%d]=Name:%s==>\t", t.Next.No, t.Next.Data)
		t = t.Next
		if t.Next == nil {
			fmt.Println()
			break
		}
	}
}

//26.2.6 双向链表的逆序打印方法
func (head *BidirectionalChainNode) ReversePrint() {
	t := head
	if t.Next == nil {
		fmt.Println("该链表为空")
		return
	}
	//定位到最后节点（这里逆序打印只是为了演示，为双向环形链表做知识储备）
	for {
		if t.Next == nil {
			break
		}
		t = t.Next
	}
	for {
		fmt.Printf("BidirectionalChainNode[%d]=Name:%s\t==>\t", t.No, t.Data)
		t = t.Pre
		if t.Pre == nil {
			fmt.Println()
			break
		}
	}
}

//26.2.7 双向链表的按序号插入
func (head *BidirectionalChainNode) Insert(newBCN *BidirectionalChainNode) {
	t := head
	for {
		if t.Next == nil {
			break
		}
		if t.Next.No > newBCN.No {
			break
		} else if t.Next.No == newBCN.No {
			fmt.Println("已经存在这个No了：", t.Next)
			return
		}
		t = t.Next
	}
	newBCN.Next = t.Next
	newBCN.Pre = t
	if t.Next != nil {
		t.Next.Pre = newBCN
	}
	t.Next = newBCN
}

//26.2.8 双向链表的删除
func (head *BidirectionalChainNode) Del(no int) {
	t := head
	for {
		if t.Next == nil {
			fmt.Println("不存在这个No")
			return
		}
		if t.Next.No == no {
			break
		}
		t = t.Next
	}
	t.Next = t.Next.Next
	if t.Next != nil {
		t.Next.Pre = t
	}
}

//测试BidirectionalChainNode
func BidirectionalChainNodeTest() {
	//1.创建一个头结点
	head := &BidirectionalChainNode{}

	//2.创建一个新的SingleChan
	sc1 := &BidirectionalChainNode{
		No:   1,
		Data: "小明",
	}
	sc2 := &BidirectionalChainNode{
		No:   2,
		Data: "小红",
	}
	sc3 := &BidirectionalChainNode{
		No:   3,
		Data: "小兰",
	}
	head.Append(sc1)
	head.Append(sc2)
	head.Append(sc3)
	head.Print()
	head.Del(3)
	head.Print()
}

//26.3.1 单向环形链表
type UnidirectionalCircularChainNote struct {
	No   int
	Data string
	Next *UnidirectionalCircularChainNote
}

func (head *UnidirectionalCircularChainNote) Append(newUCCN *UnidirectionalCircularChainNote) {
	//判断是否为第一个节点
	if head.Next == nil {
		head.No = newUCCN.No
		head.Data = newUCCN.Data
		head.Next = head
		return
	}
	t := head
	for {
		if t.Next == head {
			break
		}
		t = t.Next
	}
	t.Next = newUCCN
	newUCCN.Next = head
}

func (head *UnidirectionalCircularChainNote) Print() {
	t := head
	if t.Next == nil {
		fmt.Println("该链表为空")
		return
	}
	for {
		fmt.Printf("UnidirectionalCircularChainNote[%d]=Name:%s==>\t", t.No, t.Data)
		if t.Next == head {
			break
		}
		t = t.Next
	}
	fmt.Println()
}

func (head *UnidirectionalCircularChainNote) Del(no int) {
	t1 := head
	t2 := head
	if t1.Next == nil {
		fmt.Println("链表不存在")
		return
	}
	//只有一个节点的情况
	if t1.Next == head {
		t1.Next = nil
		return
	}

	for {
		if t2.Next == head {
			break
		}
		t2 = t2.Next
	}

	for {
		if t1.Next == head {
			if t1.No == no {
				t2.Next = t1.Next
			} else {
				fmt.Println("该No不存在：", no)
			}
			break
		}
		if t1.No == no {
			if t1 == head {
				*head = *head.Next
				break
			}
			t2.Next = t1.Next
			break
		}
		t1 = t1.Next
		t2 = t2.Next
	}

}

// var (
// 	goroutineComputLock sync.Mutex //同步.互斥  可以但不推荐（效率低）
// )

// //25 闭环队列
// type CircleQueue struct {
// 	MaxSize int //5
// 	Arrary  []int
// 	Head    int
// 	Tail    int
// }

// //放入
// func (c *CircleQueue) Push(val int) (err error) {
// 	if c.IsFull() {
// 		return errors.New("queue full")
// 	}
// 	c.Arrary[c.Tail] = val
// 	c.Tail = (c.Tail + 1) % (c.MaxSize + 1)
// 	return nil
// }

// //取出
// func (c *CircleQueue) Pop() (val int, err error) {
// 	if c.IsEmpty() {
// 		return -1, errors.New("queue empty")
// 	}
// 	val = c.Arrary[c.Head]
// 	c.Head = (c.Head + 1) % (c.MaxSize + 1)
// 	return
// }

// //显示
// func (c *CircleQueue) Show() {
// 	size := c.Size()
// 	if size == 0 {
// 		fmt.Println("队列为空")
// 		return
// 	}
// 	t := c.Head
// 	for i := 0; i < size; i++ {
// 		fmt.Printf("Arrary[%d]=%d\n", t, c.Arrary[t])
// 		t = (t + 1) % c.MaxSize
// 	}
// }

// //判断环形队列是否为满
// func (c *CircleQueue) IsFull() bool {
// 	return (c.Tail+1)%c.MaxSize == c.Head
// }

// //判断环形队列是否为空
// func (c *CircleQueue) IsEmpty() bool {
// 	return c.Tail == c.Head
// }

// //判断环形队列有多少个元素
// func (c *CircleQueue) Size() int {
// 	return (c.Tail + c.MaxSize - c.Head) % c.MaxSize
// }

// func Server(i int, c *CircleQueue) {
// 	var randseed = time.Now().UnixNano()
// 	for {
// 		randseed++
// 		rand.Seed(randseed)
// 		randNum := rand.Intn(10) + 1
// 		time.Sleep(time.Second * time.Duration(randNum))
// 		goroutineComputLock.Lock()
// 		val, err := c.Pop()
// 		goroutineComputLock.Unlock()
// 		if err == nil {
// 			fmt.Printf("%d号协程正在服务%d号客户...\n", i, val)
// 		} else {
// 			fmt.Println("err=", err)
// 		}
// 	}
// }

// func main() {
// 	c := &CircleQueue{
// 		MaxSize: 5,
// 	}
// 	c.Arrary = make([]int, c.MaxSize+1)
// 	for i := 1; i < 3; i++ {
// 		go Server(i, c)
// 	}
// 	var randseed = time.Now().UnixNano()
// 	for i := 1; ; i++ {
// 		randseed++
// 		rand.Seed(randseed)
// 		randNum := rand.Intn(5) + 1
// 		time.Sleep(time.Second * time.Duration(randNum))
// 		goroutineComputLock.Lock()
// 		err := c.Push(i)
// 		goroutineComputLock.Unlock()
// 		if err == nil {
// 			fmt.Printf("%d号客户在排队...\n", i)
// 		} else {
// 			fmt.Println("err=", err)
// 		}
// 	}
// }

//27.1 选择排序（快于冒泡排序）
func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		maxIndex := i
		max := arr[i]
		for j := i + 1; j < len(arr); j++ {
			if max > arr[j] {
				max = arr[j]
				maxIndex = j
			}
		}
		//交换
		if maxIndex != i {
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		}
	}
}

func SelectionSortTest() {
	arr := []int{2, 4, 3, 6, 1, 100, 20}
	fmt.Println("排序前：", arr)
	SelectionSort(arr)
	fmt.Println("排序后:", arr)
}

//27.2 插入排序（快于选择排序）
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		NodeVal := arr[i]
		NodeIndx := i - 1
		for NodeIndx >= 0 && arr[NodeIndx] > NodeVal {
			arr[NodeIndx+1] = arr[NodeIndx]
			NodeIndx--
		}
		if NodeIndx+1 != i {
			arr[NodeIndx+1] = NodeVal
		}
	}
}

func InsertionSortTest() {
	arr := []int{2, 4, 3, -6, 100, 52}
	fmt.Println("排序前：", arr)
	InsertionSort(arr)
	fmt.Println("排序后：", arr)
}

//27.3 快速排序(最快的)
func QuickSort(leftIndex int, rightIndex int, arr []int) {
	l := leftIndex
	r := rightIndex
	pivot := arr[(l+r)/2]
	t := 0
	for l < r {
		for arr[l] < pivot {
			l++
		}
		for arr[r] > pivot {
			r--
		}
		if l >= r {
			break
		}
		t = arr[l]
		arr[l] = arr[r]
		arr[r] = t
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	if leftIndex < r {
		QuickSort(leftIndex, r, arr)
	}
	if rightIndex > l {
		QuickSort(l, rightIndex, arr)
	}
}

func QuickSortTest() {
	arr := []int{2, 4, 3, -6, 100, 52}
	fmt.Println("排序前：", arr)
	QuickSort(0, len(arr)-1, arr)
	fmt.Println("排序后：", arr)
}

//28 栈(先入后出的有序列表)
//插入和取出均在同一端：栈顶  固定的一端：栈底
//28.1 切片模拟栈
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
func (s *StackStruct) Print() {
	if s.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	for i := s.Top; i > -1; i-- {
		fmt.Printf("arr[%v]=%v\n", i, s.Arr[i])
	}
}

func StackTest() {
	s := &StackStruct{
		Cap: 10,
		Top: -1,
	}
	s.Arr = make([]int, s.Cap)
	s.Push(555)
	s.Push(569)
	s.Push(24)
	s.Print()
	s.Pop()
	v, err := s.Pop()
	if err == nil {
		fmt.Println("v=", v)
	} else {
		fmt.Println("err=", err)
	}
	s.Pop()
	s.Print()

}
