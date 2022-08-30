package work

import ()

type Component interface {
	StartAll()
	StartOne()
}
type ComponentDefault struct {
	works []Work
}

func (c*ComponentDefault) StartAll(){}
func (c*ComponentDefault) StartOne(){}

func GetComponent() Component{
	return &ComponentDefault{}
}