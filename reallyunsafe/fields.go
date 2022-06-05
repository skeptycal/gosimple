package reallyunsafe

import "unsafe"

type Fake struct {
	Public  int
	private int
}

func NewFake(pub, pri int) *Fake {
	return &Fake{Public: pub, private: pri}
}

type Sneaky struct {
	Public  int
	Private int
}

func Recast(s Fake) Sneaky {
	return *(*Sneaky)(unsafe.Pointer(&s))
}
