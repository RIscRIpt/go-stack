package stack

type stack struct {
	top  *element
	size uint
}

type element struct {
	value interface{}
	next  *element
}

func New() *stack {
	return &stack{}
}

func (s *stack) Size() uint { return s.size }

func (s *stack) Push(value interface{}) {
	s.top = &element{
		value: value,
		next:  s.top,
	}
	s.size++
}

func (s *stack) Top() (value interface{}) {
	if s.size > 0 {
		value = s.top.value
	}
	return
}

func (s *stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
	}
	return
}

func (s *stack) Clear() {
	//TODO: Make sure gc cleans up memory pointed by s.top!
	s.top = nil
	s.size = 0
}
