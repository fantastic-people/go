package main
import (
	"fmt"
	"regexp"
)
const text="my email is sad@bac.com"
func main() {
	re:=regexp.MustCompile('([a-zA-Z0-9])+@([a-zA-Z0-9])+\.([a-zA-Z0-9])+')
	match:=re.FindString(text)
	fmt.Println(match)
}