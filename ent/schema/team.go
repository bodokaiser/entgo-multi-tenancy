package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/bodokaiser/entgo-multi-tenancy/ent/privacy"
	"github.com/bodokaiser/entgo-multi-tenancy/ent/schema/rule"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the Team.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).Through("members", Member.Type),
	}
}

// Policy defines the policies of the Team.
func (Team) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.FilterTeam(),
			privacy.AlwaysAllowRule(),
		},
		Mutation: privacy.MutationPolicy{
			rule.DenyIfNoUser(),
			privacy.AlwaysAllowRule(),
		},
	}
}
