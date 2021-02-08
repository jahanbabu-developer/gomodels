package main
import"fmt"
func main(){
	s:= struct {
		first string
		friends map[string]int
		favdrink []string
	}{
		first: "jahan",
		friends: map[string]int{
			"mohamed":569,
			"sivas":655,
			"raja":566,
		},
		 favdrink: []string{
			 "tea",
		 "milkshake",
		},
	}
	  fmt.Println(s.first)
	  fmt.Println(s.friends)
	  fmt.Println(s.favdrink)
	  for k, v:=range s.friends{
		  fmt.Println(k,v)
	  }		 
	  for i,val:=range s.favdrink{
		  fmt.Println(i,val)
	  }
	
}
