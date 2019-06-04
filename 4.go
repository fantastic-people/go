package main
import (
	"fmt"
	"encoding/json"
	"os"
	"bufio"
	"io"
)
type Monster struct{
	Name string
	Age int
	Birthday string
	sal float64
	Skill string
}
func testStruct(){
	monster:=Monster{
		Name:"老牛",
		Age:14,
		Birthday:"2018-5-98",
		sal:500.0,
		Skill:"老牛圈",
	}
	data,err:=json.Marshal(&monster)
	if err!=nil{
		fmt.Printf("monster序列化错误 err=%v\n",err)
	}
	fmt.Printf("monster序列化后=%v\n",string(data))
}
func testMap(){
	var a map[string]interface{}
	a=make(map[string]interface{})
	a["name"]="女孩儿"
	a["age"]=30
	a["address"]="北京"
	data,err:=json.Marshal(&a)
	if err!=nil{
		fmt.Printf("monster序列化错误 err=%v\n",err)
	}
	fmt.Printf("monster序列化后=%v\n",string(data))
}
func testFloat64(){
	var num float64=734.235
	data,err:=json.Marshal(&num)
	if err!=nil{
		fmt.Printf("monster序列化错误 err=%v\n",err)
	}
	fmt.Printf("monster序列化后=%v\n",string(data))
}
func testSlice(){
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1=make(map[string]interface{})
	m1["name"]="孙悟空"
	m1["age"]=30
	m1["address"]="北京"
	slice=append(slice,m1)

	var m2 map[string]interface{}
	m2=make(map[string]interface{})
	m2["name"]="的撒旦"
	m2["age"]=555
	m2["address"]="北京"
	slice=append(slice,m2)
	data,err:=json.Marshal(&slice)
	if err!=nil{
		fmt.Printf("monster序列化错误 err=%v\n",err)
	}
	fmt.Printf("monster序列化后=%v\n",string(data))
	
}
type CharCount struct{
	ChCount int 
	NumCount int 
	SpaceCount int 
	OtherCount int 
}
func main(){
	// testStruct()
	// testMap()
	// testFloat64()
	fileName:="e:/abc.txt"
	file,err:=os.Open(fileName)
	if err!=nil{
		fmt.Printf("Open file err=%v\n",err)
		return 
	}
	defer file.Close()
	var count CharCount
	reader:=bufio.NewReader(file)
	for{
		str,err:=reader.ReadString('\n')
		if err ==io.EOF{
			break
		}
		for _,v:=range str {
			switch  {
			case v>='a' && v<='z':
				fallthrough
			case v>='A' && v<='z':
				count.ChCount++
			case v==' ' || v=='\t':
				count.SpaceCount++
			case v>='0' && v<='9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Printf("字符的个数=%v,数字的个数=%v，空格的个数=%v,其他字符的个数=%v\n",count.ChCount,count.NumCount,count.SpaceCount,count.OtherCount)
}