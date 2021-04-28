package main

import (
	"fmt"
	"runtime"
)

type T struct{
	name string //name of the object
	value int //value of the object

}
type BT struct{
left BT//left child
right BT//right child
}
func getstr(){
var a string="123"
for i:= range 10{
a+='c'
}
return a
}

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		a:=1
		b:=2
		c:=a+b
		fmt.Printf("%s.\n", os)
	}
}
