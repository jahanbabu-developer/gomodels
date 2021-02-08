package main
import (
	"fmt"
)
type dog int
var x dog
func main(){
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x=42
	fmt.Println(x)
}