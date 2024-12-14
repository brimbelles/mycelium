package main

type propertyType uint

const (
	text propertyType = iota
	integer
	double
	boolean
)

type propertyData struct {
	propertyDataType propertyType
	stringData       string
	integerData      int64
	doubleData       float64
	booleanData      bool
}

type Property struct {
	Key   string
	Value propertyData
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
