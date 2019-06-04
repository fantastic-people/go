package main
import(
	"fmt"
	
)
// func main() {
//     ch := make(chan int, 3)
//     ch <- 1
//     ch <- 2
//     ch <- 3

//     fmt.Println(<-ch) // 1
//     fmt.Println(<-ch) // 2
//     fmt.Println(<-ch) // 3
// }
func main() {
    // ch := make(chan int, 3)
    // ch <- 1
    // ch <- 2
    // ch <- 3

    // for v := range ch {
    //     fmt.Println(v)
	// }
	ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3
for v := range ch {
    fmt.Println(v)
    if len(ch) <= 0 { // 如果现有数据量为0，跳出循环
        break
    }
}
}