package main

import (
	. "gc"
)

func main() {
	obj1 := &Object{}
	obj2 := &Object{}
	obj3 := &Object{}

	obj1.References = []*Object{obj2}
	obj2.References = []*Object{obj3}

	gc := NewGarbageCollector(obj1, []*Object{obj1, obj2, obj3})
	gc.Collect()
}
