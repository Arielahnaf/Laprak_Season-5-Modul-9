package main

import "fmt"

func main() {
	var atmin int
	var hasil bool
	fmt.Scan(&atmin)
	// if atmin > 0 && abi%2 == 0 {
	// 	hasil = true
	// }
	hasil = atmin > 0 && atmin%2 == 0
	fmt.Print(hasil)
}
