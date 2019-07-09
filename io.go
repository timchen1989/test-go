package main
import "fmt"
func main(){
	var a int
	var b int
	fmt.Print("please input some number :")
	fmt.Scanln(&a, &b)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a+b)
	// var c int
	// fmt.Scanln("please input some number : ", &c)

}	
