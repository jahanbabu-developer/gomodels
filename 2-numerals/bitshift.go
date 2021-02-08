package main

import(
	"fmt"
)
var a int=20
func main(){
	fmt.Printf("%d\t%b\t%#x\n",a,a,a)
	b:=a<<1
	fmt.Printf("%d\t%b\t%#x",b,b,b)
}