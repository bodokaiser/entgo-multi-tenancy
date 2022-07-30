package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("owner").
			Default(false),
		field.Bool("admin").
			Default(false),
		field.Int("team_id"),
		field.Int("user_id"),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("team", Team.Type).
			Required().
			Unique().
			Field("team_id"),
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id"),
	}
}

// Annotations of the Member.
func (Member) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("team_id", "user_id"),
	}
}
