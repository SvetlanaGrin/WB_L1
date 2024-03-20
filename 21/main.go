package main

import "fmt"

func main() {

	adapter := NewAdapter(&Adaptee{})

	req := adapter.Request()

	fmt.Println(req)

}
