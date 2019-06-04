package main
import (
    "fmt"
    "runtime"
)

var quit chan int = make(chan int)

// func loop() {
//     for i := 0; i < 100; i++ { //为了观察，跑多些
//         fmt.Printf("%d ", i)
//     }
//     quit <- 0
// }
func loop() {
    for i := 0; i < 10; i++ {
        runtime.Gosched() // 显式地让出CPU时间给其他goroutine
        fmt.Printf("%d ", i)
    }
    quit <- 0
}
func main() {
    runtime.GOMAXPROCS(2) // 最多使用2个核

    go loop()
    go loop()

    for i := 0; i < 2; i++ {
        <- quit
    }
}