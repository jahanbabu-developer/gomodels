package main

import(
	"fmt"
	"sync"
	"runtime"
)
func main(){

fmt.Println("begin cpu",runtime.NumCPU())
fmt.Println("begin gs",runtime.NumGoroutine())

var wg sync.WaitGroup
wg.Add(2)

go func(){
	fmt.Println("hello from thing one")
    wg.Done()
	}()
go func(){
	fmt.Println("hello from thing two")
	wg.Done()
}()
fmt.Println("mid cpu",runtime.NumCPU())
fmt.Println("mid gs",runtime.NumGoroutine())
wg.Wait()
fmt.Println("About to exit")
fmt.Println("end cpu",runtime.NumCPU())
fmt.Println("end gs",runtime.NumGoroutine())
}
