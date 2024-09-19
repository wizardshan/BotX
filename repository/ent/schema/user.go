package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("hash_id").Default(""),
		field.String("mobile").Default(""),
		//field.Int("age").Optional().Positive(),
		//field.Bool("active").Default(true),
		//
		//field.String("nickname").NotEmpty().MinLen(2).MaxLen(10),
		//field.String("password").Sensitive(),
		//field.Enum("level").GoType(user.Level(0)),
	}
}

//func (User) Mixin() []ent.Mixin {
//	return []ent.Mixin{
//		mixin.Time{},
//	}
//}
//
//// Edges of the User.
//func (User) Edges() []ent.Edge {
//	return nil
//}
