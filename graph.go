package main

type Property struct {
	Key   string
	Value string
}

type Node struct {
	Label      string
	Properties []Property
}

type Relationship struct {
	From       Node
	To         Node
	Properties []Property
}
