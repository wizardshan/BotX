# BotX

Fields
All fields are required by default, and can be set to optional using the Optional method.

func (User) Fields() []ent.Field {
    return []ent.Field{
        field.String("mobile"),
        field.Int("age"),
    }
}
db.User.Create().SetAge(0).Save(context.Background()) 
����ent: missing required field "User.mobile"
������field.String("mobile").Optional()
������field.String("mobile").Default("")