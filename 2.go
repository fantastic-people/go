package main 
import (
	"fmt"
	"time"
	"math/rand"
)
func CreatNum(p *int)  {
	rand.Seed(time.Now().UnixNano())
	var num  int
	for{
		num=rand.Intn(10000)
		if num>=1000{
			break
		}
	}
	*p=num
}
func GetNum(s[]int,num int){
	s[0]=num/1000
	s[1]=num%1000
	s[2]=num%100
	s[3]=num%10
}
func OnGame(randSlice [] int){
	var num int 
	keySlice:=make([]int, 4)
	for{
		for{
			fmt.Println("请输入一个4位数： ")
			fmt.Scan(&num)
			if 999<num && 10000>num {
				break
			}
			fmt.Println("您输入的不符合要求，请重新输入！")
		}
		GetNum(keySlice,num)
		n:=0
		for index := 0; index < 4; index++ {
			if keySlice[index] > randSlice[index]{
				fmt.Printf("第%d位大了一点！\n",index+1)
				//fmt.Println("第",index+1,"位大了一点\n")
			}else if keySlice[index] < randSlice[index]{
				fmt.Printf("第%d位小了一点！\n",index+1)
				//fmt.Println("第",index+1,"位小了一点\n")
			}else{
				//fmt.Println("第",index+1,"猜对了！\n")
				fmt.Printf("第%d位猜对了！！\n",index+1)
				n++;
			}
		}
		if n==4  {
			//fmt.Println("全部猜对了！！！\n")
			fmt.Printf("全部猜对了！！！\n")
			break
		}
	}
}
func main1(){
	var randNum int
	CreatNum(&randNum)
	randSlice:=make([]int, 4)
	GetNum(randSlice,randNum)
	OnGame(randSlice)	
}