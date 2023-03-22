package mongo

import architecture "github.com/rcondongfox/golang-architecture"

type MongoDB map[int]architecture.Person

func (m MongoDB) Save(n int, p architecture.Person) {
	m[n] = p
}

func (m MongoDB) Retrieve(n int) architecture.Person {
	return m[n]
}