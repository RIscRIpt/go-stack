package stack

import "testing"

func TestNew(t *testing.T) {
	if New() == nil {
		t.Fatalf("New() returned nil!")
	}
}

func TestSize0(t *testing.T) {
	if New().Size() != 0 {
		t.Fatalf("New stack size is not 0!")
	}
}

func TestPushTop(t *testing.T) {
	s := New()
	s.Push("element")
	if s.Top() != "element" {
		t.Fatalf("Top is %v", s.Top())
	}
}

func TestSize2(t *testing.T) {
	s := New()
	s.Push(true)
	s.Push(3.14)
	if s.Size() != 2 {
		t.Fatalf("Stack size is not valid! (%v)", s.Size())
	}
}

func TestPushPop(t *testing.T) {
	test := []interface{}{1, "b", 'c', 4.5}
	s := New()
	for _, x := range test {
		s.Push(x)
	}
	for i := len(test) - 1; i >= 0; i-- {
		if s.Pop() != test[i] {
			t.FailNow()
		}
	}
}

func TestPushClear(t *testing.T) {
	s := New()
	s.Push(3.14)
	s.Push(2.72)
	s.Clear()
	if s.Size() != 0 {
		t.Fatalf("Clear() didn't clear the stack!")
	}
}
