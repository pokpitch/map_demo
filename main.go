package main

import "fmt"

func main() {
	//var m map[string]string
	//m = make(map[string]string)
	m := map[string]string{}

	if m == nil {
		fmt.Println("it's nil")
	} else {
		fmt.Println("it's not nil")
	}

	m["a"] = "apple"
	s := m["a"]

	fmt.Println(s)

}
