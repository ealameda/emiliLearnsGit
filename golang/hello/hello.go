package main

import "fmt"

func main() {
	fmt.Println(Hello("anything"))
}

func Hello(name string) string {
	return "Hello, " + name
}
