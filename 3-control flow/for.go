package main

import (
"fmt"
)
func main(){
birth:=1998
	for {
		if birth>=2021{
			break
		}
		fmt.Println(birth)
		birth++
	}
}