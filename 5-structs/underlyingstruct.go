package main
import(
	"fmt"
)
type vehicle struct{
	doors int
	color string
}
type truck struct{
vehicle
fourwheeler bool

}
type sedan struct{
vehicle
luxury bool
}
func main(){
	t:=truck{
	    vehicle : vehicle{
		doors:2,
		color:"black",
},
fourwheeler:true,
	}
	s:=sedan{
		vehicle : vehicle{
			doors:4,
			color:"silver",
},
luxury:false,
	}
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.doors)
	fmt.Println(s.doors)

}
