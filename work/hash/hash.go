package main

import (
	"fmt"
	"os"
)

type Emp struct {
	Id   int
	Name string
	Next *Emp
}

func (e *Emp) ShowMe() {
	fmt.Printf("链表%d找到该雇员%d", e.Id%7, e.Id)
}

//这里的EmpLink 不带表头，即第一个结点就存放雇员
type EmpLink struct {
	Head *Emp
}

//1.添加员工的方法,保证添加时编号从小到大
func (el *EmpLink) Insert(emp *Emp) {
	cur := el.Head     //这是辅助指针
	var pre *Emp = nil //辅助指针 pre 在cur前面
	//如果当前的EmpLink就是一个空链表
	if cur == nil {
		el.Head = emp //完成
		return
	}
	//如果不是一个空链表，给emp找到对应的位置并插入
	//让cur和emp比较，然后让pre保持在cur前面
	for {
		if cur != nil {
			//比较
			if cur.Id > emp.Id {
				//找到位置
				break
			}
			pre = cur //保证同步
			cur = cur.Next
		} else {
			break
		}
	}
	//退出时，我们看下是否将emp添加到链表最后
	pre.Next = emp
	emp.Next = cur
}

//显示当前链表的信息
func (el *EmpLink) ShowLink(no int) {
	if el.Head == nil {
		fmt.Printf("链表%d为空\n", no)
		return
	}

	//变量当前的链表，并显示数据
	cur := el.Head //辅助指针
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id%d 名字=%s->\n", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}

//根据id查找对应的雇员，如果没有就返回nil
func (el *EmpLink) FindById(id int) *Emp {
	cur := el.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

func (el *EmpLink) DelById(id int) bool {
	if el.Head == nil {
		return false
	}
	if el.Head.Next == nil { //只有一个的情况
		if el.Head.Id == id {
			el.Head = nil
			return true
		} else {
			return false
		}
	}
	//如果第一个刚好就是：
	if el.Head.Id == id {
		el.Head = el.Head.Next
		return true
	}
	pre := el.Head
	cur := el.Head.Next
	for {
		if cur.Id == id {
			pre.Next = cur.Next
			return true
		}
		pre = pre.Next
		cur = cur.Next
		if cur == nil {
			return false
		}
	}

}

type HashTable struct {
	LinkArr [7]EmpLink
}

//给hashTable 编写Insert 雇员的方法
func (ht *HashTable) Insert(emp *Emp) {
	//使用散列函数，确定将该雇员添加到哪个链表
	linkNo := ht.HashFunc(emp.Id)
	//使用对应的链表添加
	ht.LinkArr[linkNo].Insert(emp) //
}

//显示hashtable的所有雇员
func (ht *HashTable) ShowAll() {
	for i := 0; i < len(ht.LinkArr); i++ {
		ht.LinkArr[i].ShowLink(i)
	}
}

//编写一个散列方法
func (ht *HashTable) HashFunc(id int) int {
	return id % 7 //得到一个值，就是对于链表的下标
}

//查找
func (ht *HashTable) FindById(id int) *Emp {
	//使用散列函数，确定该雇员应该在哪个链表
	linkNo := ht.HashFunc(id)
	return ht.LinkArr[linkNo].FindById(id)
}

func (ht *HashTable) DelById(id int) bool {
	//使用散列函数，确定该雇员应该在哪个链表
	linkNo := ht.HashFunc(id)
	return ht.LinkArr[linkNo].DelById(id)
}

func main() {
	num := 0
	id := 0
	name := ""
	var hashTable HashTable
	for {
		fmt.Println("===雇员菜单===")
		fmt.Println("1.添加雇员")
		fmt.Println("2.显示雇员")
		fmt.Println("3.查找雇员")
		fmt.Println("4.删除雇员")
		fmt.Println("5.退出")
		fmt.Println("请输入你的选择：")
		fmt.Scanln(&num)
		switch num {
		case 1:
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员name")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashTable.Insert(emp)
		case 2:
			hashTable.ShowAll()
		case 3:
			fmt.Println("输入要查找的雇员id")
			fmt.Scanln(&id)
			emp := hashTable.FindById(id)
			if emp == nil {
				fmt.Printf("id=%d的雇员不存在\n", id)
			} else {
				//编写一个方法，显示雇员信息
				emp.ShowMe()
			}
		case 4:
			fmt.Println("输入要删除的雇员id")
			fmt.Scanln(&id)
			b := hashTable.DelById(id)
			if !b {
				fmt.Println("删除失败")
			} else {
				fmt.Println("删除成功")
			}
		case 5:
			os.Exit(0)
		default:
			fmt.Println("输入错误！")
		}

	}
}
