// package main

// import (
//     "fmt"
// )

// func squares() func() int {
// 	var x int
// 	x=5
//     return func() int {
// 		for i:=0;i<x;i++{
// 			x=i+x
// 		}
//         return x
//     }
// }

// func main() {
//     f1 := squares()
//     f2 := squares()

//     fmt.Println("first call f1:", f1())
// 	fmt.Println("second call f1:", f1())
// 	fmt.Println("third call f1:", f1())
	
	

//     fmt.Println("first call f2:", f2())
// 	fmt.Println("second call f2:", f2())    
// 	fmt.Println("third call f2:", f2())
                                              
// }

// package main
 
// import(
//     "fmt"
//     "time"
// )
 
// func main(){
//     for i:=0;i<100;i++{
//         go func(){
//             fmt.Println(i)
//         }()
//     }
//     time.Sleep(1e9)
// }
// package main
 
// import "fmt"
// func make_fn() func(i int){
//     return func(i int){
//         fmt.Println(i)
//     }
// }
// func main(){
 
//     var fn [10]func(int)
 
//     for i:=0;i<len(fn);i++{
//         fn[i]=make_fn()
//     }
 
//     for i,f:=range fn{ // 闭包中使用的是这个ide指针
//         f(i)
//     }
// }
 
package main
import "fmt"

func ExFunc(n int) func() {
         sum:=n
         return func () { //把匿名函数作为值赋给变量a (Go 不允许函数嵌套。
                      sum++             //然而你可以利用匿名函数实现函数嵌套)
                  fmt.Println(sum+1) //调用本函数外的变量
         } //这里没有()匿名函数不会马上执行
}

func main() {
         myFunc:=ExFunc(10)
         myFunc()     // 11
         myAnotherFunc:=ExFunc(20)
         myAnotherFunc()    //21
         myFunc()       //11
         myAnotherFunc()   //21
}