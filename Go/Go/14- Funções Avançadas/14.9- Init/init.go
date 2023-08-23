package main

import "fmt"


// Esta função roda antes da main
func init() {
	fmt.Println("oieeee")
}


func main() {
	fmt.Println("oiiiiii")
}