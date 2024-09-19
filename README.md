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
报错：ent: missing required field "User.mobile"
不报错：field.String("mobile").Optional()
不报错：field.String("mobile").Default("")


TDD：测试驱动开发（Test-Driven Development）
DDD：领域驱动设计（Domain Drive Design）