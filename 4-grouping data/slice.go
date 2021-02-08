package main
import"fmt"
func main(){
	a:=[]int{2,3,4,5,9,12,14,1,2,15}
	for i,v:=range a{
		fmt.Println(i,v)
		
	}
	fmt.Printf("%T\n",a)
}