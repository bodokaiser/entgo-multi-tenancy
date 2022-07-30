package auth

import (
	"context"

	"github.com/bodokaiser/entgo-multi-tenancy/ent"
)

type ctxKeyUser struct{}

func WithUser(ctx context.Context, u *ent.User) context.Context {
	return context.WithValue(ctx, ctxKeyUser{}, u)
}

func UserFrom(ctx context.Context) *ent.User {
	u, _ := ctx.Value(ctxKeyUser{}).(*ent.User)
	return u
}
