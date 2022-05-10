package stack

// Stack is a LIFO sequential access data structure.
// Access O(n)
// Search O(n)
//* Insert O(1)
//* Delete O(1)
type (
	Stack[E comparable, S ~[]*E] interface {
		Push(value E)
		Pop() E
		Peek() E
		Contains(value E) bool
		Len() int
		Cap() int
		// Grow(n float64)
	}

	stack[E comparable, S ~[]*E] struct {
		buf S
	}
)

func New[E comparable, S ~[]*E](size int) Stack[E, S] {
	return &stack[E, S]{buf: make(S, size, size*2)}
}

func (s *stack[E, S]) Peek() E  { return *s.buf[s.max()] }
func (s *stack[E, S]) Len() int { return len(s.buf) }
func (s *stack[E, S]) Cap() int { return cap(s.buf) }
func (s *stack[E, S]) max() int { return len(s.buf) - 1 }

func (s *stack[E, S]) Push(value E) {
	if s.capCheck() {
		s.grow(0)
	}
	s.buf = append(s.buf, &value)
	// s.buf[len(s.buf)+1] = &value
}

const (
	minStackGrow                 = 5   // TODO test various ...
	minStackShrink               = .5  // TODO test various ...
	defaultStackGrowMultiplier   = 1.1 // TODO test various ...
	defaultStackShrinkMultiplier = 0.8
)

func (s *stack[E, S]) Pop() E {
	if s.shrinkCheck() {
		s.shrink(0)
	}
	// TODO add shrink check
	r := s.Peek()
	s.buf = s.buf[:len(s.buf)-1]
	return r
}

func (s *stack[E, S]) Contains(value E) bool {
	for _, e := range s.buf {
		if *e == value {
			return true
		}
	}
	return false
}

// Grow reallocates the stack to a higher capacity
// based on the multiplier, e.g. a multiplier of
// 1.2 will grow the capacity by 120%.
// Using a value of 0 will automate the process.
//
// Negative values, values less than 1, or values
// greater than 10 will be normalized to achieve
//  1.1 < multiplier < 10
func (s *stack[E, S]) grow(n float64) {
	size := s.newGrowSize(n)

	b := make(S, len(s.buf), size)
	copy(b, s.buf)
	s.buf = b
}

// shrink reallocates the stack to a lower capacity
// based on the multiplier, e.g. a multiplier of
// 0.6 will shrink the capacity to 60% or original size.
// Using a value of 0 will automate the process.
//
// Negative values or values greater than 1 will be
// normalized to achieve
//  0 < multiplier < 1
func (s *stack[E, S]) shrink(n float64) {
	size := s.newShrinkSize(n)
	if size < s.Len() {
		return
	}

	// s.buf = s.buf[:size]
	b := make(S, len(s.buf), size)
	copy(b, s.buf)
	s.buf = b
}

// newShrinkSize returns the new target capacity of the stack.
func (s *stack[E, S]) newShrinkSize(multiplier float64) int {
	for multiplier >= 1 {
		multiplier *= 0.1
	}

	switch m := multiplier; {
	case m == 0:
		multiplier = defaultStackShrinkMultiplier
	case m < 0:
		return s.newShrinkSize(-multiplier)
	case m > defaultStackShrinkMultiplier:
		multiplier += defaultStackShrinkMultiplier
	}

	return int(multiplier * float64(cap(s.buf)))
}

// newGrowSize returns the new target capacity of the stack.
func (s *stack[E, S]) newGrowSize(multiplier float64) int {

	for multiplier >= 10 {
		multiplier *= 0.1
	}

	switch m := multiplier; {
	case m == 0:
		multiplier = defaultStackGrowMultiplier
	case m < 0:
		return s.newGrowSize(-multiplier)
	case m < defaultStackGrowMultiplier:
		multiplier += defaultStackGrowMultiplier
	}

	return int(multiplier * float64(cap(s.buf)))
}

// capCheck returns true if the stack size is within
// the default grow size of the capacity. Time to grow!
func (s *stack[E, S]) capCheck() bool {
	return int(float64(len(s.buf))*defaultStackGrowMultiplier) > cap(s.buf)
}

// capCheck returns true if the stack size is within
// the default shrink size of the capacity. Time to shrink!
func (s *stack[E, S]) shrinkCheck() bool {
	return int(float64(len(s.buf))*defaultStackShrinkMultiplier) < cap(s.buf)
}
