package main

import (
	"botx/controller"
	"botx/pkg/sms"
	"botx/repository"
	"botx/repository/ent"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/bbs?charset=utf8mb4&parseTime=true"
	db, err := ent.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	//level := user.Level(0)
	//u, err := db.User.Create().SetAge(1).Save(context.Background())

	engine := gin.New()
	repoUser := repository.NewUser(db)
	ctrUser := controller.NewUser(repoUser)
	engine.GET("/user/:hashID", controller.Wrapper(ctrUser.Show))
	engine.GET("/user", controller.Wrapper(ctrUser.One))
	engine.GET("/user/login", controller.Wrapper(ctrUser.Login))
	engine.GET("/user/register", controller.Wrapper(ctrUser.Register))

	smsClient := sms.NewClient()
	ctrSms := controller.NewSms(smsClient)
	engine.GET("/sms/captcha", controller.Wrapper(ctrSms.Captcha))

	engine.Run()
}
