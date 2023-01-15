package main
import "fmt"

func sum2(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum2(1,2)

	sum2(1,2,3)

	nums := []int{1,2,3,4}
	sum2(append(nums,-100)...)

}