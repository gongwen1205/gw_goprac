package main
//add log2
//add log5
//add log4
//add log3
import (
	"fmt"
	"runtime"
"math/rand"
)

type T struct{
	name string //name of the object
	value int //value of the object

}
type BT struct{
left BT//left child
right BT//right child
}
//add struct
type DL struct{
prev DL
next DL
}
//add log
func getstr(){
var a string="123"
for i:= range 10{
a+='c'
}
return a
}
func getfloat(){
var a float=1.0
a=a*3
return a
}
//add bool log
func getbool(){
var a bool=1
return a
}
func getnu() int{
a:=0
a=rand.Intn(10)
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
