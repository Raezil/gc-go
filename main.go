package main

import "fmt"

const (
	White = iota
	Gray
	Black
)

type Object struct {
	color      int
	references []*Object
}
type GarbageCollector struct {
	heap    []*Object
	graySet []*Object
	root    *Object
}

func NewGarbageCollector(root *Object, heap []*Object) *GarbageCollector {
	return &GarbageCollector{
		root:    root,
		heap:    heap,
		graySet: []*Object{},
	}
}

func (gc *GarbageCollector) InitializeWhiteObjects() {
	for _, obj := range gc.heap {
		obj.color = White
	}
}

func (gc *GarbageCollector) MarkGray(obj *Object) {
	obj.color = Gray
	gc.graySet = append(gc.graySet, obj)
}

func (gc *GarbageCollector) Collect() {
	gc.InitializeWhiteObjects()
	gc.MarkGray(gc.root)
	for len(gc.graySet) > 0 {
		currentObject := gc.graySet[len(gc.graySet)-1]
		gc.graySet = gc.graySet[:len(gc.graySet)-1]
		for _, reference := range currentObject.references {
			if reference.color == White {
				gc.MarkGray(reference)
			}
		}
		currentObject.color = Black
	}
	for _, obj := range gc.heap {
		if obj.color == White {
			fmt.Println("Collecting unreachable object")
		}
	}
}

func main() {
	obj1 := &Object{}
	obj2 := &Object{}
	obj3 := &Object{}

	obj1.references = []*Object{obj2}
	obj2.references = []*Object{obj3}

	gc := NewGarbageCollector(obj1, []*Object{obj1, obj2, obj3})
	gc.Collect()
}
