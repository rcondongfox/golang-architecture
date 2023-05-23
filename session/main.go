package session

import "context"

type stringKey string
type intKey int

var userId stringKey
var admin intKey

func SetUserId(ctx context.Context, Id int) context.Context {
	return context.WithValue(ctx, userId, Id)
}

func SetAdmin(ctx context.Context, isAdmin bool) context.Context {
	return context.WithValue(ctx, admin, isAdmin)
}

func GetUserId(ctx context.Context) int {
	if uId := ctx.Value(userId); uId != nil {
		if i, ok := uId.(int); ok {
			return i
		}
	} 
	return 0
}

func GetAdmin(ctx context.Context) bool {
	if v := ctx.Value(admin); v != nil {
		if i, ok := v.(bool); ok {
			return i
		}
	} 
	return false
} 