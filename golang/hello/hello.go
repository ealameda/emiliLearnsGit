package main

import "fmt"

func main() {
	fmt.Println(Hello("anything"))
}

func Hello(name string) string {
	if name == "" {
		return "Hello, World"
	}
	return "Hello, " + name
}
