package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    ch := make(chan struct{}, 2) 

    var l sync.Mutex


    go func() {
        fmt.Println("groutine2: 等待解锁")
        l.Lock()
        defer l.Unlock()
        fmt.Println("goroutine2: 哈哈，我锁定了")
        ch <- struct{}{}
    }()
    go func() {
        l.Lock()
        defer l.Unlock()
        fmt.Println("goroutine1: 我会锁定大概 2s")
         time.Sleep(time.Second * 2)
        fmt.Println("goroutine1: 我解锁了，你们去抢吧")
        ch <- struct{}{}
    }()
    // 等待 goroutine 执行结束
    for i := 0; i < 2; i++ {
        <-ch
    }
}