package main
import (
	"testing"
	"fmt"
)
func TestPrint(t *testing.T)  {
	res:=Print1to20()
	fmt.Println("hey")
	if res!=210{
		t.Errorf("Wrong result of Print1to20")
	}
}