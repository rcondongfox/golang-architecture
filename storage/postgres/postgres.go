package postgres

import architecture "github.com/rcondongfox/golang-architecture"

type PostgresDB map[int]architecture.Person

func (pg PostgresDB) Save(n int, p architecture.Person) {
	pg[n] = p
}

func (pg PostgresDB) Retrieve(n int) architecture.Person {
	return pg[n]
}