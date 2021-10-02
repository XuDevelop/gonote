package note

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// func main(){
// 	fmtPrintingVerb()
// }

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

//1.3 随机数生成
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

//1.4 匿名函数
func AnonymousFunction() {
	fmt.Println("1.4 匿名函数")
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

//3.1 基本数据类型和string的转换
func BasicDataTypeAndStringConversion() {
	//3.1.1 基本数据类型转成string
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

	//3.1.2 string转成基本数据类型
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

//3.2 指针pointer
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

//3.3 值类型和引用类型

//值类型：基本数据类型int系列，float系列，bool，string、数组和结构体
//变量直接存储值，内存通常在栈中分配
//（内存分栈区和堆区。）
//栈区：值类型数据,通常在栈区
//堆区：引用类型，通常在堆区分配空间
//引用类型：指针、slice切片、map、管道chan、interface。。。
//变量存储是一个地址，这个地址对应的空间才是真正存储数据（值），内存通常在堆区上分配，当没有任何变量引用这个地址时，
//改地址对应的数据空间就成为了一个垃圾，由GC来回收

//3.5 字符串的常用函数
func FunctionsInStrings() {
	fmt.Println("3.5.1 统计字符串的长度，按字节 len(str) : func len(v Type) int ")
	//数组：v中元素的数量  数组指针：*v中元素的数量（v为nil时panic）
	//切片、映射：v中元素的数量；若v为nil，len（v）即为零
	//字符串：v中字节的数量    通道：通道缓存中队列（未读取）元素的数量；若v为nil，len（v）即为零
	//var a int =123456
	var b string = "hello"
	//a1:=len(a)
	b1 := len(b)
	//fmt.Println("var a int =123456 ; a1:=len(a) ; 输出：",a1)
	fmt.Println("var b string =\"hello\" ; b1:=len(b) ; 输出：", b1)

	fmt.Println("3.5.2 字符串遍历，同时处理有中文的问题 r：=[]rune(str)")
	str2 := "hello你"
	for i := 0; i < len(str2); i++ {
		fmt.Println("字符=", str2[i])
	} //打印的全是数字
	str3 := []rune(str2) //字符串遍历，同时处理有中文的问题 r：=[]rune(str)
	for i := 0; i < len(str3); i++ {
		fmt.Printf("字符=%c\n", str2[i])
	} //打印出字符
	fmt.Println(str3)

	fmt.Println("3.5.3 字符串转整数：func Atoi(s string)(i int,err error)")
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换失败", err)
	} else {
		fmt.Println("转换成功")
		fmt.Printf("Type:%T,value:%v\n", n, n)
	}

	fmt.Println("3.5.4 整数转字符串：str=strconv.Itoa(123)")
	str4 := strconv.Itoa(123)
	fmt.Printf("Type:%T,value:%v\n", str4, str4)

	fmt.Println("3.5.5 字符串转[]byte：var byte=[]byte(\"hello word\")")
	var bytes = []byte("hello word")
	fmt.Printf("Type=%T,char=%c\n", bytes, bytes)

	fmt.Println("3.5.6 []byte转字符串: str=string([]byte{11,22,33})")
	str6 := string([]byte{11, 22, 33})
	fmt.Printf("Type=%T,int=%v\n", str6, str6)

	fmt.Println("3.5.7 查找字符串是否在指定的字符串中：strings.Contains(\"baby\",\"honey\")")
	//func Contains(s,suber string)bool  如果没有指定的字符则输出false
	c := strings.Contains("abc", "ad")
	fmt.Println(c)

	fmt.Println("3.5.8 统计一个字符里有几个指定的字符串：strings.Count(\"abcc\",\"c\")") //区分大小写的
	d := strings.Count("abssBss", "b")
	fmt.Println(d)

	fmt.Println("3.5.9 区分和不区分大小写的字符串比较")
	fmt.Println(strings.EqualFold("abc", "ABC")) //不区分大小写
	fmt.Println("abc" == "ABC")                  //区分大小写

	fmt.Println("3.5.10 返回子串在字符第一次出现的index值，如果没有返回-1  strings.Index(\"NMnhh abc\",\"abc\")")
	//strings.Index("NMnhh abc","abc")
	fmt.Println(strings.Index("NMnhh_abc", "abc")) //6

	fmt.Println("3.5.11 返回子串在字符最后一次出现的index值，如果没有返回-1   strings.LastIndex(\"NMnhh abc\",\"abc\")")
	fmt.Println(strings.LastIndex("NMnhh_abcabcabcabcabc", "abc")) //18

	fmt.Println("3.5.12 将指定的子串替换成另一个子串：strings.Replace(\"go go hello\",\"go\",\"go语言\",n)")
	// fmt.Println(strings.Replace("go go hello","go","go语言",n))   n 可以指定你希望替换几个 n=-1是全部替换
	fmt.Println(strings.Replace("go go hello", "go", "go语言", -1))

	fmt.Println("3.5.13 按照指定的某个字符为分割标识，将字符串拆分成字符串数组  strings.Split(\"hello word,ok\",\",\")")
	zf := strings.Split("hello word,ok", ",")
	for i := 0; i < len(zf); i++ {
		fmt.Println(zf[i])
	}
	fmt.Println(strings.Split("hello word,ok", ","))

	fmt.Println("3.5.14 将字符串的字母进行大小写的转换  strings.ToLower(\"Go\")")
	//strings.ToLower("Go")//go
	//strings.ToUpper("Go")//GO
	fmt.Println(strings.ToLower("Go"))
	fmt.Println(strings.ToUpper("Go"))

	fmt.Println("3.5.15 将字符串左右两边的空格去掉：")
	q := strings.TrimSpace("  aa ss   aa aa   a a a  ")
	fmt.Println(q)

	fmt.Println("3.5.16 将字符串左右两边指定的字符去掉：")
	w := strings.Trim("!aa ss!!aa aa!   a !!a a !!", "!")
	fmt.Println(w)

	fmt.Println("3.5.17 将字符串左边指定的字符去掉：")
	e := strings.TrimLeft("!aa ss!!aa aa!   a !!a a !! ", "!")
	fmt.Println(e)

	fmt.Println("3.5.18 将字符串右边指定的字符去掉：")
	r := strings.TrimRight("!aa ss!!aa aa!   a !!a a !!", "!")
	fmt.Println(r)

	fmt.Println("3.5.19 判断字符串是否以指定的字符串开头")
	t := strings.HasPrefix("abc,wwes22", "abc")
	fmt.Println(t)

	fmt.Println("3.5.20 判断字符串是否以指定的字符串结束")
	y := strings.HasSuffix("sdwdewdxwccmm", "m")
	fmt.Println(y)
}

//3.6 时间和日期函数
func TimeAndData() {
	fmt.Println("3.6 时间和日期函数")
	//3.6.1完整版
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

	//3.6.2时间的常量
	// const(
	// 	Nanosecond Duration=1//纳秒        //Duration是int64的别名
	// 	Microsecond =1000*Nanosecond//微秒
	// 	Millisecond =1000*Microsecond//毫秒
	// 	Second=1000*Millisecond//秒
	// 	Minute=60*Second//分钟
	// 	Hour=60*Minute//小时
	// )

	//3.6.3 休眠
	//func Sleep(d Duration)
	//time.Seelp(100*time.Millisecond)//休眠100毫秒
	//案例：
	for i := 1; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(5 * time.Second)
	}

	//3.6.4 获取当前unix时间戳和unixnano时间戳   //1970-01-01 00:00

}

//3.7 new(Type)
func FunctionNew() {
	fmt.Println("3.7 new(Type)")
	a := new(int) //传出地址跟var a *int差不多
	fmt.Printf("Type=%T,valeu=%v", a, a)
}

//3.8 错误捕捉延迟处理  //出错后，会输出错误但是整个程序继续执行
func SearchDaly() {
	fmt.Println("3.8 错误捕捉延迟处理  //出错后，会输出错误但是整个程序继续执行")
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

//3.9 自定义错误
func CustomError() {
	fmt.Println("3.9 自定义错误")
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

//4  数组
func Array() {
	//4.1 申明数组
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

	//4.2 For...range遍历数组
	for index, value := range arr {
		fmt.Println("第", index, "个数组的值是：", value)
	}

	//4.3 声明切片
	//4.3.1 引用数组
	slice1 := make([]int, 10)
	s := slice1
	slice1[0] = 666
	fmt.Println(s[0])
	//4.3.2 make([]type,len,cap(optional))    len是定义该切片的长度 cap是容量必须>=len
	var slice2 = make([]int, 3) //or slice2:=make([]int,3)
	slice2[0] = 99
	fmt.Println(slice2[0])

	//4.4 append的用法
	//4.4.1 用法一：
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)
	//4.4.2 用法二：
	slice2 = append(slice2, 3)
	fmt.Println(slice2)

	//4.5 copy()各种情况
	copy(slice1, []int{11, 12, 13, 14})
	fmt.Println(slice1)
	var a = []int{1, 2}
	copy(a, []int{9, 9, 9, 9, 9})
	fmt.Println(a)
}

//5 map,key是无序的，map是引用类型
func MapAndKey() {
	//5.1 声明map，make（），和直接声明，赋值
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

	//5.2 删除元素 delete()
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

//6 创建结构体实例
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

//7 工厂模式编程
//7.1 工厂模式结构体
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

//8 继承
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

//9 接口 interface
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

//10 类型断言
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

//11 文件操作
func FileOperation() {
	filePath := "E:/develop/gonote/file.txt"
	// 11.1 打开与关闭
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

	//11.2 带缓冲区的读取写文件（大文件）
	//11.2.1 带缓冲区的读取，适合大文件的读取，需要手动打开与关闭
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

	//11.2.2 带缓冲区的写入，适合大文件，需要手动打开与关闭
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

	//11.3.1 一次性读取文件（小文件）
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("一次性读取文件失败！", err)
		return
	}
	fmt.Println("一次性读取到的：", string(content))

	//11.3.2 一次性写入文件（小文件）
	content = []byte("pig" + "fangpig")
	err = ioutil.WriteFile(filePath, content, 0666) //这种写入是直接覆盖整个文件
	if err != nil {
		fmt.Println("一次性写入文件失败！", err)
		return
	}

	//11.4 判断文件是否存在
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

	//11.5 拷贝文件
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
