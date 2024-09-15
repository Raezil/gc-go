package gc

import "testing"

func TestGarbageCollector_Collect(t *testing.T) {
	gc := NewGarbageCollector(nil, nil)
	err := gc.Collect()
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestGarbageCollector_InitializeWhiteObjects(t *testing.T) {
	gc := NewGarbageCollector(nil, nil)
	gc.InitializeWhiteObjects()
	for _, obj := range gc.heap {
		if obj.color != White {
			t.Errorf("expected object color to be white, got %d", obj.color)
		}
	}
}

func TestGarbageCollector_MarkGray(t *testing.T) {
	gc := NewGarbageCollector(nil, nil)
	obj := &Object{}
	gc.MarkGray(obj)
	if obj.color != Gray {
		t.Errorf("expected object color to be gray, got %d", obj.color)
	}
	if len(gc.graySet) != 1 {
		t.Errorf("expected gray set length to be 1, got %d", len(gc.graySet))
	}
}

func TestGarbageCollector_NewGarbageCollector(t *testing.T) {
	root := &Object{}
	heap := []*Object{root}
	gc := NewGarbageCollector(root, heap)
	if gc.root != root {
		t.Errorf("expected root to be set, got nil")
	}
	if len(gc.heap) != 1 {
		t.Errorf("expected heap length to be 1, got %d", len(gc.heap))
	}
	if len(gc.graySet) != 0 {
		t.Errorf("expected gray set length to be 0, got %d", len(gc.graySet))
	}
}
