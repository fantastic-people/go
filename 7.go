package main
import(
	"fmt"
	"os"
)
func main(){
	dir,_:=os.Getwd()
	fmt.Println("当前目录为：",dir)
	path:=os.Getenv("GOPATH")
	fmt.Println("当前GOPATH目录为：",path)
}