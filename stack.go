package stack

type Stack struct {
	top  *element
	size uint
}

type element struct {
	value interface{}
	next  *element
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Size() uint { return s.size }

func (s *Stack) Push(value interface{}) {
	s.top = &element{
		value: value,
		next:  s.top,
	}
	s.size++
}

func (s *Stack) Top() (value interface{}) {
	if s.top != nil {
		value = s.top.value
	}
	return
}

func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
	}
	return
}

func (s *Stack) Rop() interface{} {
	s.Pop()
	return s.Top()
}

func (s *Stack) Clear() {
	//TODO: Make sure gc cleans up memory pointed by s.top!
	s.top = nil
	s.size = 0
}
