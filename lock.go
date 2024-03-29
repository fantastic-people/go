package main
import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var a = 0

    // 启动 100 个协程，需要足够大
    var lock sync.Mutex
    for i := 0; i < 100; i++ {
        go func(idx int) {
            lock.Lock()
			
            a += 1
			fmt.Printf("goroutine %d, a=%d\n", idx, a)
			defer lock.Unlock()
		}(i)
		
    }

    // 等待 1s 结束主程序
    // 确保所有协程执行完
    time.Sleep(time.Second)
}


