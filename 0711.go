package main
import ("fmt")
func main(){
	for i := 0;i<5;i++{
		fmt.Println(i)
	}
	// fmt.Println(i)
	fmt.Println(split(64))
}
func split(sum int) (x, y int){
	x = sum * 4/9
	// y = sum - x
	return
}