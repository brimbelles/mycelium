package main

import "fmt"

func main() {

	n1 := Node{"n1", nil}
	n2 := Node{"n2", nil}

	p1 := StringProperty("p1-key", "p1-value")
	p2 := IntegerProperty("p2-key", 2)
	p3 := DoubleProperty("p3-key", 3.0)
	p4 := BooleanProperty("p4-key", true)
	r12, err := NewRelationship(&n1, &n2, []Property{p1, p2, p3, p4})
	if err != nil {
		fmt.Println(err)
	}

	g1 := Graph{[]Node{n1, n2}, []Relationship{r12}}

	fmt.Println(g1)

	s1 := LabeledNodes(&g1, "n2")
	fmt.Println(s1)
}
