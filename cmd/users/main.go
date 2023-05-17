package main

import (
	"fmt"

	architecture "github.com/rcondongfox/golang-architecture"
	Db "github.com/rcondongfox/golang-architecture/storage/mongo"
)

func main() {
	storage := Db.Db{}

	p1 := architecture.User{Name: "John"}
	p2 := architecture.User{Name: "Jane"}

	architecture.Put(storage, 1, p1)
	architecture.Put(storage, 2, p2)

	fmt.Println(architecture.Get(storage, 1))
	fmt.Println(architecture.Get(storage, 2))


}