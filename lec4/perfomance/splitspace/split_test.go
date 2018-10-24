package splitspace

// go test -benchmem -bench .

import (
	"testing"
)

var (
	in = `The second are its interfaces, which provides runtime polymorphism. Interfaces are a class of types and provide a limited form of 
	structural typing in the otherwise nominal type system of Go. An object which is of an interface type is also of another type, much like 
	C++ objects being simultaneously of a base and derived class. Go interfaces were designed after protocols from the Smalltalk programming
	 language. Multiple sources use the term duck typing when describing Go interfaces. Although the term duck typing is not precisely 
	 defined and therefore not wrong, it usually implies that type conformance is not statically checked. Since conformance to a Go interface 
	 is checked statically by the Go compiler (except when performing a type assertion), the Go authors prefer the term structural typing.`
)

func BenchmarkTrim(b *testing.B) {
	for i := 0; i < b.N; i++ {
		splitTrim(in)
	}
}

func BenchmarkRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tokenize(in)
	}
}

func BenchmarkRegexpGlobal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tokenizeGlobal(in)
	}
}
