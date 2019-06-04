package main
import (
	"fmt"
)
func getCharCount(str string) map[byte]int{
	buf:=make(map[byte]int)
	for index := 0; index < len(str); index++ {
		_, ok :=buf[str[index]]
		if ok{
			buf[str[index]]+=1
		}else{
			buf[str[index]]=1
		}
	}
	return buf
}
func main()  {
	fmt.Printf("请输入字符串：")
	var str string
	fmt.Scanln(&str)
	buf:=getCharCount(str)
	for key,value:=range buf{
		fmt.Printf("%c出现%d次数\n",key,value)
	}
}