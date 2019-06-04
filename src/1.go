
package main 
import (
	"fmt"
	"strings"
)
func main(){
	fmt.Println(strings.ContainsAny("widnn","w&d"))
	fmt.Println(strings.ContainsAny("widnn","w&d&n"))
}