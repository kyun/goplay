package main

import "fmt"

func main() {
	r := methods.rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
}
