package main
import (	
	"fmt"
	"flag"
)

func main1(){
	//fmt.Printf("hello")
	ok:=flag.Bool("ok",false,"id ok")
	fmt.Println("OK:",*ok)
}

