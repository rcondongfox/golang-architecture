package main

import (
	"context"
	"fmt"

	"github.com/rcondongfox/golang-architecture/session"
)

func main() {
	
	ctx := context.Background()
	ctx =  session.SetUserId(ctx, 123)
	ctx = session.SetAdmin(ctx, true)
	
	uId := session.GetUserId(ctx)
	fmt.Printf("User id: %d\n", uId)
	isAdmin := session.GetAdmin(ctx)
	fmt.Printf("User %d is admin: %v\n", session.GetUserId(ctx), isAdmin)

}