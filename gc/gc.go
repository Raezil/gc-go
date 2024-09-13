package gc

import "fmt"

const (
	White = iota
	Gray
	Black
)

// Object represents an object in the heap
type Object struct {
	color      int
	References []*Object
}

// GarbageCollector represents a garbage collector that collects unreachable objects
type GarbageCollector struct {
	heap    []*Object
	graySet []*Object
	root    *Object
}

// NewGarbageCollector creates a new garbage collector
func NewGarbageCollector(root *Object, heap []*Object) *GarbageCollector {
	return &GarbageCollector{
		root:    root,
		heap:    heap,
		graySet: []*Object{},
	}
}

// InitializeWhiteObjects initializes all objects in the heap to white
func (gc *GarbageCollector) InitializeWhiteObjects() {
	for _, obj := range gc.heap {
		obj.color = White
	}
}

// MarkGray marks an object as gray and adds it to the gray set
func (gc *GarbageCollector) MarkGray(obj *Object) {
	obj.color = Gray
	gc.graySet = append(gc.graySet, obj)
}

// Collect collects unreachable objects
func (gc *GarbageCollector) Collect() error {
	if gc.root == nil {
		return fmt.Errorf("root object is nil")
	}
	gc.InitializeWhiteObjects()
	gc.MarkGray(gc.root)
	for len(gc.graySet) > 0 {
		currentObject := gc.graySet[len(gc.graySet)-1]
		gc.graySet = gc.graySet[:len(gc.graySet)-1]
		for _, reference := range currentObject.References {
			if reference.color == White {
				gc.MarkGray(reference)
			}
		}
		currentObject.color = Black
	}
	var newHeap []*Object
	for _, obj := range gc.heap {
		if obj.color == White {
			fmt.Println("Collecting unreachable object")
		} else {
			newHeap = append(newHeap, obj)
		}
	}
	gc.heap = newHeap
	return nil
}
