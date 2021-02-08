package main
import"fmt"
type person struct{
first string	
}
type human interface{
speak()
}
func(*person)speak(){
	fmt.Println("hello")
}
func saySomething(h human){
	h.speak()
}
func main(){
	p1:=person{
		first:"james",
	}	
	saySomething(&p1)
	p1.speak()
}
