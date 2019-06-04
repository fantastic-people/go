package main 
import (
	"fmt"
)
func main()  {
	
	defer fmt.Println(1)
	defer fmt.Println(12)
	fmt.Println(123)
	fmt.Println(1234)
	fmt.Println(12345)
	defer fmt.Println(123456)
	for index := 0; index < 5; index++ {
		defer fmt.Println(index)
	}
}