package hind

// hind provides a "len(a)-1" function, hind.Hind.  It is slower than
// if it was implemented as part of the Go language or by using
// generics.  That said, it serves as a proof-of-concept to see if
// this would be a useful addition to the language.

// If we had generics, I'd call this end(), highest(), or ultimate().

import "reflect"

// G returns the index of the last item in any (generic) array, slice, or
// string.  It is not as efficient as "len(a)-1", but better conveys
// intention.
// It serves as a placeholder for when generics makes this easy and
// efficient.
// Typical usage: hint.G(x)
func G(a interface{}) int {
	// Code comes from a suggestion by C Banning <clbanning@gmail.com>
	v := reflect.ValueOf(a)
	switch v.Kind() {
	case reflect.Array, reflect.Slice, reflect.String:
		return v.Len() - 1
	}
	return 0
}

// Str returns the highest index of a string. The compiler should
// inline this, making it as efficient as "len(a)-1" (allowing the
// optimizer to eliminate the -1 when needed).
// It serves as a placeholder for when generics makes this easy and
// efficient.
// Typical usage: hint.S(s)
func S(a string) int {
	return len(a) - 1
}
