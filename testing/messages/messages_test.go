package messages

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Parallel()
	got := Greet("Gopher")
	expect := "Hello, Gopher\n"
	if got != expect {
		t.Errorf("Did not get expected result")
	}
}

func TestDepart(t *testing.T) {
	t.Parallel()
	got := depart("Gopher")
	expect := "Goodbye, Gopher\n"
	if got != expect {
		t.Errorf("Did not get expected result")
	}

}

func TestGreetTableDriven(t *testing.T) {
	t.Parallel()
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "Gopher", expect: "Hello, Gopher\n"},
		{input: "", expect: "Hello, \n"},
	}
	for _, s := range scenarios {
		t.Run(s.input, func(t *testing.T) {
			got := Greet(s.input)
			if got != s.expect {
				t.Errorf("Error for %v", s.input)
			}
		})
	}
}

func BenchmarkSHA1(b *testing.B) {
	data := []byte("uihequrgfahwriukfghelruihgehgosr")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha1.Sum(data)
	}
}

func BenchmarkSHA256(b *testing.B) {
	data := []byte("uihequrgfahwriukfghelruihgehgosr")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha256.Sum256(data)
	}
}

func BenchmarkSHA512(b *testing.B) {
	data := []byte("uihequrgfahwriukfghelruihgehgosr")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha512.Sum512(data)
	}
}

func BenchmarkSHA512Allocations(b *testing.B) {
	data := []byte("uihequrgfahwriukfghelruihgehgosr")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		h := sha512.New()
		h.Sum(data)
	}
}
