package main

import "errors"

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

func StringProperty(key string, value string) Property {
	return Property{key, propertyData{propertyDataType: text, stringData: value}}
}

func IntegerProperty(key string, value int64) Property {
	return Property{key, propertyData{propertyDataType: integer, integerData: value}}
}

func DoubleProperty(key string, value float64) Property {
	return Property{key, propertyData{propertyDataType: double, doubleData: value}}
}

func BooleanProperty(key string, value bool) Property {
	return Property{key, propertyData{propertyDataType: boolean, booleanData: value}}
}

type Node struct {
	Label      string
	Properties []Property
}

type Relationship struct {
	From       *Node
	To         *Node
	Properties []Property
}

func NewRelationship(from, to *Node, properties []Property) (Relationship, error) {
	rel := Relationship{from, to, properties}
	if from == nil {
		return rel, errors.New("Origin node is nil")
	}
	if to == nil {
		return rel, errors.New("Destination node is nil")
	}
	return rel, nil
}

type Graph struct {
	nodes         []Node
	relationships []Relationship
}
