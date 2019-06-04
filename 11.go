// package main
// import(
// 	"fmt"
// 	"time"
// )
// func Add(i ,j int)  {
// 	z:= i+j
// 	fmt.Println(z)
// }
// func main(){
// 	for i:=0;i<10;i++ {
// 		go Add(i,i)
// 	}
// 	// time.Sleep(time.Second)
// }
package main 
import "fmt" 
import "sync" 
import "runtime" 
var counter int = 0 
func Count(lock *sync.Mutex) { 
 lock.Lock() 
 counter++ 
 fmt.Println(counter) 
 lock.Unlock() 
} 
func main() { 
 lock := &sync.Mutex{} 
 for i := 0; i < 10; i++ { 
 go Count(lock) 
 } 
for { 
 lock.Lock() 
 c := counter 
 lock.Unlock() 
 runtime.Gosched() 
 if c >= 10 { 
 break 
 } 
 } 
} 
