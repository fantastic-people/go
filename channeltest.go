package main

import (
    "fmt"
    // "time"
)
var echo chan string
var receive chan string

// 定义goroutine 1 
func Echo() {
    // time.Sleep(1*time.Second)
    echo <- "咖啡色的羊驼"
    fmt.Println(3)
}

// 定义goroutine 2
func Receive() {
    temp := <- echo // 阻塞等待echo的通道的返回
    receive <- temp
    fmt.Println(2)
}


func main() {
    echo = make(chan string)
    receive = make(chan string)

    go Echo()
    go Receive()
    ch := make(chan string, 3) // 创建了缓冲区为3的通道
    
    //=========
    fmt.Println(len(ch))  // 长度计算
    fmt.Println(cap(ch))   // 容量计算
    // getStr := <-receive   // 接收goroutine 2的返回
    <- ch
     fmt.Println(1)
}