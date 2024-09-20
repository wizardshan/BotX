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

HashID 
db => domain 直接赋值
request => domain 1、查询需要检查 hashID =>id 2、创建需要 id => hashID

Password
db => domain 直接赋值
request => domain 需要检查

Level
db => domain 需要获得
request => domain 需要检查

