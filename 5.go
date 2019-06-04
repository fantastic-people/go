package main
import(
	"fmt"
	"util"
)
func main(){
	// fmt.Println("utils.go num-=um1")
	fmt.Println("utils.go num=",util.Num1)
	var n1 float64=1.2
	var n2 float64=2.3
	var operator byte='+'
	result:=util.Cal(n1,n2,operator)
	fmt.Println("result=",result)
	 n1 =6
	 n2 =3
	 operator='*'
	result=util.Cal(n1,n2,operator)
	fmt.Println("result=",result)

}