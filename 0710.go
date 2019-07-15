package main
import ("fmt"
"math")
func main(){
	fmt.Println("pi : ", math.Pi)
	fmt.Println("x+y : ", add(1, 2))
	fmt.Println("x+y+z : ", add2(2,3,4))
	x, y := swap("a", "b")
	fmt.Println(x, y)
	var a = 1
	// a := 2
	b := 2
	fmt.Println(b)
	fmt.Println(a)
}
func add(x int, y int) int {
	return x + y
}
func add2(x, y, z int ) int {
	return x + y + z
}
func swap(x, y string) (string, string){
	return y, x
}

