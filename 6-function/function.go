package main
import"fmt"
func main(){
m:=foo()
fmt.Println(m)
n,s:=bar()
fmt.Println(n,s)
}
func foo() int{
 return 12
} 
func bar() (int,string){
	return 24,"hours"

}