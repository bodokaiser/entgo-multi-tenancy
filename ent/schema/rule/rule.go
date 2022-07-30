package rule

import (
	"context"

	"github.com/bodokaiser/entgo-multi-tenancy/auth"
	"github.com/bodokaiser/entgo-multi-tenancy/ent"
	"github.com/bodokaiser/entgo-multi-tenancy/ent/member"
	"github.com/bodokaiser/entgo-multi-tenancy/ent/predicate"
	"github.com/bodokaiser/entgo-multi-tenancy/ent/privacy"
	"github.com/bodokaiser/entgo-multi-tenancy/ent/user"
)

// DenyIfNoUser is a rule that returns a dency decision if a user is missing from the context.
func DenyIfNoUser() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		if auth.UserFrom(ctx) == nil {
			return privacy.Denyf("user missing from context")
		}
		return privacy.Skip
	})
}

// FilterTeam is a rule that filters out teams not associated to the user by a membership.
func FilterTeam() privacy.QueryMutationRule {
	type TeamFilter interface {
		WhereHasMembersWith(...predicate.Member)
	}

	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		u := auth.UserFrom(ctx)
		if u == nil {
			return privacy.Denyf("user missing from context")
		}

		tf, ok := f.(*ent.TeamFilter)
		if !ok {
			return privacy.Denyf("incompatible filter type %T", f)
		}

		tf.WhereHasMembersWith(member.HasUserWith(user.ID(u.ID)))

		return privacy.Skip
	})
}
