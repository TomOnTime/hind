package hind

import "testing"

func TestHind_strarray(t *testing.T) {
	var strarray = []string{
		"hello",
		"out",
		"there",
		"in",
		"the",
		"world"}

	g1 := G(strarray)
	w := 5
	if g1 != w {
		t.Errorf("G() on strarray: got %v, wanted %v", g1, w)
	}
}

func TestHind_intarray(t *testing.T) {
	var intarray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	w := 9

	g1 := G(intarray)
	if g1 != w {
		t.Errorf("G() on intarray: got %v, wanted %v", g1, w)
	}

}

func TestHind_slice(t *testing.T) {
	var intarray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var s = intarray[2:5]
	w := 2

	g1 := G(s)
	if g1 != w {
		t.Errorf("G() on slice: got %v, wanted %v", g1, w)
	}

}

func TestHind_str(t *testing.T) {
	var str = "hello, world"
	w := 11

	g1 := G(str)
	if g1 != w {
		t.Errorf("G() on string: got %v, wanted %v", g1, w)
	}

	g2 := S(str)
	if g2 != w {
		t.Errorf("S() on string: got %v, wanted %v", g2, w)
	}
}
