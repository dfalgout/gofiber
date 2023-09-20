package auth

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

func Add(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, userCtxKey, user)
}

func Get(ctx context.Context) string {
	return ctx.Value(userCtxKey).(string)
}

func GetUser(c *fiber.Ctx) (*string, bool) {
	ctx := c.UserContext()
	if ctx == nil {
		return nil, false
	}
	found := Get(ctx)
	return &found, true
}
