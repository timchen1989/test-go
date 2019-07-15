package main
import ("fmt")
func main(){
	test8()
}
func test1(){
	fmt.Println("test")
	var i int
	var f float64
	var b bool
	var s string
	var n rune
	fmt.Println("=", i, f, b, s, n, "=")
} 
func test2() {
	var i int = 12
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(i, f, u)
	a := 42
	b := float64(a)
	c := uint(b)
	fmt.Println(a, b, c)
	fmt.Printf("a Type : %T\n", a)
	fmt.Printf("b Type : %T\n", b)
	fmt.Printf("c Type : %T\n", c)
	fmt.Printf("s %T\n", "Test")
}
func test3() {
	const hello = "hello"
	const world = "world"
	fmt.Println(hello, Pi, world)
	// hello = "w"
	const Truth = true
	fmt.Println("Go rules?", Truth)
}
const Pi = 3.14
func test4() {
	sum := 0
	for i := 0; i<= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
func test5() {
	sum := 1
	for sum <=1000 {
		sum += sum
	fmt.Println(sum)
	}
	fmt.Println(sum)
}
func test6(b bool) {
	if b {
		fmt.Println("if condition result : True")
	} else {
		fmt.Println("if condition result : False")
	}
}
func test7() {
	if b := false ; b {
		fmt.Println("if condition result : True")
	} else {
		fmt.Println("if condition result : False")
	}
	// fmt.Println(b)
}
func test8() {
	for i :=0 ; i <=9 ; i++ {
		for j := 0 ; j <= 9 ; j++ {
			fmt.Print(i , " x ", j, " = ", i*j, " ")
		} 
		fmt.Println()
	}
}