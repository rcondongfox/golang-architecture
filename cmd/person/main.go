package main

import (
	"fmt"

	architecture "github.com/rcondongfox/golang-architecture"
	mongo "github.com/rcondongfox/golang-architecture/storage/mongo"
	postgres "github.com/rcondongfox/golang-architecture/storage/postgres"
)

func main() {
	dbMongo := mongo.MongoDB{}
	dbp := postgres.PostgresDB{}
	
	p1 := architecture.Person{
		FirstName: "Richy",
	}
	
	p2 := architecture.Person{
		FirstName: "Hichy",
	}
	
	p3 := architecture.Person{
		FirstName: "Jim",
	}

	p4 := architecture.Person{
		FirstName: "Sarah",
	}
	
	p5 := architecture.Person{
		FirstName: "Tony",
	}
	
	p6 := architecture.Person{
		FirstName: "Erik",
	}


	mongoPersonService := architecture.NewPersonService(dbMongo)
	postgresPersonService := architecture.NewPersonService(dbp)
	
	//STORE IN MONGO
	architecture.Put(dbMongo, 1, p1)
	architecture.Put(dbMongo, 2, p2)
	architecture.Put(dbMongo, 3, p3)

	fmt.Println("Database Entry 1, using architecture.Get() :")
	fmt.Println(architecture.Get(dbMongo, 1))
	fmt.Println("Database Entry 2, using architecture.Get() :")
	fmt.Println(architecture.Get(dbMongo, 2))
	fmt.Println("Database Entry 3, using mongoPersonService.Get() :")
	fmt.Println(mongoPersonService.Get(3))
	fmt.Println("Database Entry 4, using mongoPersonService.Get() :")
	fmt.Println(mongoPersonService.Get(4))

	//STORE IN POSTGRES
	architecture.Put(dbp, 4, p4)
	architecture.Put(dbp, 5, p5)
	architecture.Put(dbp, 6, p6)

	fmt.Println("Database Entry 4, using architecture.Get() :")
	fmt.Println(architecture.Get(dbp, 4))
	fmt.Println("Database Entry 5, using architecture.Get() :")
	fmt.Println(architecture.Get(dbp, 5))
	fmt.Println("Database Entry 6, using postgresPersonService.Get() :")
	fmt.Println(postgresPersonService.Get(6))
	fmt.Println("Database Entry 7, using postgresPersonService.Get() :")
	fmt.Println(postgresPersonService.Get(7))
}