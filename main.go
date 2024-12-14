package main

import "fmt"

func main() {

	n1 := Node{"n1", nil}
	n2 := Node{"n2", nil}

	p1 := StringProperty("p1-key", "p1-value")
	p2 := IntegerProperty("p2-key", 2)
	p3 := DoubleProperty("p3-key", 3.0)
	p4 := BooleanProperty("p4-key", true)
	r12 := Relationship{n1, n2, []Property{p1, p2, p3, p4}}

	fmt.Println(r12)
}
