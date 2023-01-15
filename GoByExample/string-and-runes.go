package main
import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "Yes我爱慕课网!" // UTF-8
	fmt.Println(len(s))
	fmt.Println(s) // "Yes我爱慕课网!"

	for i := 0; i <len(s); i++{
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
}