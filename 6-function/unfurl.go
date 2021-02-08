package main
import"fmt"
func main(){
	
i:=[]int {1,2,3,4,5,6,7,8,9,10}
n:=foo(i ...)
fmt.Println(n)

j:=[]int {1,2,3,4,5,6,7,8,9}
q:=bar(j)
fmt.Println(q)
		}
		
    func foo(x ...int) int{
	total:=0
	for _,v:=range x{
	total +=v
	}
	return total
}
	func bar(x []int) int{
		total:=0
		for _,v:=range x{
		total +=v
		}
		return total 
}