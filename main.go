package main

import "fmt"

func main() {
	s := SayHello("Foo")
	fmt.Println(s)
}

func SayHello(s string) string {
	return fmt.Sprintf("Hello wrong %s!", s)
}