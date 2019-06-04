// package main 
// import "fmt" 
// func Count(ch chan int) { 
//  ch <- 1 
//  fmt.Println("Counting") 
// } 
// func main() { 
//  chs := make([]chan int, 10) 
//  for i := 0; i < 10; i++ { 
//  chs[i] = make(chan int) 
//  go Count(chs[i]) 
//  } 
//  for _, ch := range(chs) { 
//  <-ch 
//  } 
// } 
package main 
import "fmt" 
var complete chan int = make(chan int)

func loop() {
    for i := 0; i < 10; i++ {
        fmt.Printf("%d ", i)
    }

    complete <- 1 // 执行完毕了，发个消息
}


func main() {
    go loop()
    <- complete // 直到线程跑完, 取到消息. main在此阻塞住
}
// func main() {
//     ch := make(chan int)
//     <- ch // 阻塞main goroutine, 信道c被锁
// }