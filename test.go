package main
import ("fmt" 
		"time"
		"math/rand"
	)
func main(){
	fmt.Println("Time: ", time.Now())
	fmt.Println("Hello, world")
	fmt.Print("請輸入數字: ")
	var n1 int
	fmt.Scanln(&n1)
	fmt.Println("您輸入的數字為:", n1)
	fmt.Println(rand.Intn(10))
}