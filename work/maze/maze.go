package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Res struct {
	mazeMap *[8][7]int
	pass    int
}

func GetRandDirection() (x int, y int) {
	var randseed = time.Now().UnixNano()
	randseed++
	rand.Seed(randseed)
	randNum := rand.Int31n(3)
	switch randNum {
	case 0:
		return 1, 0
	case 1:
		return -1, 0
	case 2:
		return 0, 1
	case 3:
		return 0, -1
	}
	return
}

//完成老鼠找路
//mazeMap *[8][7]int：地图，保证是同一个地图，使用引用
//i,j 表示对地图的哪个点进行测试
func SetWay(mazeMap *[8][7]int, i int, j int, pass *int) bool {
	//分析出什么情况下，就找到出路
	//mazeMap[6][5]==2
	if mazeMap[6][5] == 2 {
		return true
	} else {
		//说明要继续找
		if mazeMap[i][j] == 0 { //如果这个点是可以探测的
			//假设这个点是可以通，但是需要探测上下左右
			//换个策略 下右上左
			mazeMap[i][j] = 2
			x, y := GetRandDirection()
			if SetWay(mazeMap, i+x, j+y, pass) { //下
				*pass++
				return true
			} else if SetWay(mazeMap, i+(-1)*x, j+(-1)*y, pass) { //右
				*pass++
				return true
			} else if SetWay(mazeMap, i+y, j+x, pass) { //上
				*pass++
				return true
			} else if SetWay(mazeMap, i+(-1)*y, j+(-1)*x, pass) { //左
				*pass++
				return true
			} else { //死路
				mazeMap[i][j] = 3
				return false
			}
		} else { //说明这个点不能探测，为1，是墙
			return false
		}
	}
}

func main() {
	//先创建一个二维数组，模拟迷宫
	//规则：
	//1.如果元素的值为1，就是墙
	//2.如果元素的值为0，是没有走过的点
	//3.如果元素的值为2，是一个通路
	//4.如果元素的值为3，是走过的点，但是走不通
	var mazeMap [8][7]int

	//先把地图的最上和最下设置为1
	for i := 0; i < 7; i++ {
		mazeMap[0][i] = 1
		mazeMap[7][i] = 1
	}

	//先把地图的最左和最右设置为1
	for i := 0; i < 8; i++ {
		mazeMap[i][0] = 1
		mazeMap[i][6] = 1
	}

	mazeMap[3][1] = 1
	mazeMap[3][2] = 1

	//输出地图
	fmt.Println("题目为：")
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(mazeMap[i][j], " ")
		}
		fmt.Println()
	}

	num := 32
	var res []*Res
	for i := 0; i < num; i++ {
		m := mazeMap
		//使用测试
		pass := 0
		SetWay(&m, 1, 1, &pass)
		r := &Res{
			mazeMap: &m,
			pass:    pass,
		}
		res = append(res, r)
	}

	bestRes := res[0]
	for i := 1; i < len(res); i++ {
		if res[i].pass < bestRes.pass {
			bestRes = res[i]
		}
	}

	//输出地图
	fmt.Println("最佳结果为：")
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(bestRes.mazeMap[i][j], " ")
		}
		fmt.Println()
	}
}
