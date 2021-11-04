package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
func main() {
	u := &UnidirectionalCircularChainNote{}
	num := 10
	for i := 1; i <= num; i++ {
		newUCCN := &UnidirectionalCircularChainNote{
			No: i,
		}
		u.Append(newUCCN)
	}
	var randseed = time.Now().UnixNano()
	randseed++
	rand.Seed(randseed)
	randNum := rand.Intn(num) + 1
	fmt.Println("k=", randNum)
	t1 := u //头结点交给t1
	for {
		if t1.Next == t1 {
			break
		}
		t2 := t1
		if randNum == 1 {
			for {
				if t2.Next == t1 {
					break
				}
				t2 = t2.Next
			}
			fmt.Printf("%d被干掉了\n", t1.No)
			t1.Del(t1.No)
			t1 = t2.Next
		} else {
			for j := 1; j < randNum-1; j++ {
				t2 = t2.Next
			}
			fmt.Printf("%d被干掉了\n", t2.Next.No)
			t1.Del(t2.Next.No)
			t1 = t2.Next
		}

	}
	fmt.Printf("%d活下来了\n", t1.No)
}
