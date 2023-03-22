package architecture

import (
	"fmt"
	"testing"
)


type MongoDB map[int]Person

func (m MongoDB) Save(n int, p Person) {
	m[n] = p
}

func (m MongoDB) Retrieve(n int) Person {
	return m[n]
}

func TestPut(t *testing.T) {
	mdb := MongoDB{}

	p := Person{
		FirstName: "Richy",
	}

	Put(mdb, 1, p)

	got := mdb.Retrieve(1)
	if mdb.Retrieve(1) != p {
		t.Fatalf("Want %v, got %v", p, got)
	}
}

func ExamplePut() {
	mdb := MongoDB{}

	p := Person{
		FirstName: "Richy",
	}

	Put(mdb, 1, p)
	got := mdb.Retrieve(1)
	fmt.Println(got)
	// Output: {Richy}
}