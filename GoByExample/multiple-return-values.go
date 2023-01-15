package main
import "fmt"

func vals() (int, int, string) {
	return 3, 7, "st"
}

func main() {
	a,b,cc := vals()

	fmt.Println(a)
	fmt.Println(b)
	_,_, c := vals()
	fmt.Println(c)
	fmt.Println(cc)
}