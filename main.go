package main

import "fmt"

func main() {

	n1 := Node{"n1", nil}
	n2 := Node{"n2", nil}

	r12 := Relationship{n1, n2, nil}

	fmt.Println(r12)
}
