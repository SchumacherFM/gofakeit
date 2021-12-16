package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleVerb() {
	Seed(11)
	fmt.Println(Verb())
	// Output: release
}

func ExampleFaker_Verb() {
	f := New(11)
	fmt.Println(f.Verb())
	// Output: release
}

func BenchmarkVerb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Verb()
	}
}

func ExampleVerbAction() {
	Seed(11)
	fmt.Println(VerbAction())
	// Output: close
}

func ExampleFaker_VerbAction() {
	f := New(11)
	fmt.Println(f.VerbAction())
	// Output: close
}

func BenchmarkVerbAction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VerbAction()
	}
}

func ExampleVerbLinking() {
	Seed(11)
	fmt.Println(VerbLinking())
	// Output: was
}

func ExampleFaker_VerbLinking() {
	f := New(11)
	fmt.Println(f.VerbLinking())
	// Output: was
}

func BenchmarkVerbLinking(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VerbLinking()
	}
}

func ExampleVerbHelping() {
	Seed(11)
	fmt.Println(VerbHelping())
	// Output: be
}

func ExampleFaker_VerbHelping() {
	f := New(11)
	fmt.Println(f.VerbHelping())
	// Output: be
}

func BenchmarkVerbHelping(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VerbHelping()
	}
}