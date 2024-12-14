package main

import "fmt"

func main() {

	n1 := Node{"n1", nil}
	n2 := Node{"n2", nil}
	n3 := Node{"n3", nil}
	nodes := make(map[*Node]bool)
	nodes[&n1] = true
	nodes[&n2] = true
	nodes[&n3] = true

	p1 := StringProperty("p1-key", "p1-value")
	p2 := IntegerProperty("p2-key", 2)
	p3 := DoubleProperty("p3-key", 3.0)
	p4 := BooleanProperty("p4-key", true)

	r12properties := make(map[Property]bool)
	r12properties[p1] = true
	r12properties[p2] = true
	r12properties[p3] = true
	r12properties[p4] = true

	r12, err := NewRelationship(&n1, &n2, r12properties)
	if err != nil {
		fmt.Println(err)
	}

	r23properties := make(map[Property]bool)
	r23properties[p1] = true
	r23, _ := NewRelationship(&n2, &n3, r23properties)

	relationships := make(map[*Relationship]bool)
	relationships[&r12] = true
	relationships[&r23] = true

	g1 := Graph{nodes, relationships}

	fmt.Println(g1)

	s1 := HasLabel(g1, "n2")
	for node := range s1 {
		fmt.Println(node)
	}
}
