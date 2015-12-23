# stack
Simple implementation of stack, written in [Go programming language](https://github.com/golang).

# Usage
```
import "github.com/RIscRIpt/stack"
// ...
s := stack.New()
// ...
s.Push("some value")
s.Push(true)
s.Push(3.14)
// ...
for e := s.Top(); s.Size() > 0; e = s.Pop() {
  //do something with e
}
