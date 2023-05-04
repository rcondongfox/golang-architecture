package main

import "fmt"

type user struct {
	name string
}

type mongo map[int]user

func (m mongo) save(n int, u user) {
	m[n] = u
}

func (m mongo) retrieve(n int) user {
	return m[n]
}

type hardDrive map[int]user

func (hd hardDrive) save(n int, u user) {
	hd[n] = u
}

func (hd hardDrive) retrieve(n int) user {
	return hd[n]
}

type accessor interface {
	save(n int, u user)
	retrieve(n int) user
}

func put(a accessor, n int, u user) {
	a.save(n, u)
}

func get(a accessor, n int) user {
	return a.retrieve(n)
}

func main() {
	storage := mongo{}

	p1 := user{"John"}
	p2 := user{"Jane"}

	put(storage, 1, p1)
	put(storage, 2, p2)

	fmt.Println(get(storage, 1))
	fmt.Println(get(storage, 2))


}