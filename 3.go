package main
import (
	"fmt"
	
)
type books struct{
	title string
	author string
	subject string
	book_id int
}
func main1(){
	var  Book1 books
	Book1.title="dsadasas"
	Book1.author="adasds"
	Book1.subject="dsadsas"
	Book1.book_id=1
	fmt.Printf("%s\n",Book1.title)
	fmt.Printf("%s\n",Book1.author)
	fmt.Printf("%s\n",Book1.subject)
	fmt.Printf("%d\n",Book1.book_id)
}