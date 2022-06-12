package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
)

import "github.com/gofiber/fiber/v2/middleware/recover"

func main() {
	//增加默认处理
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if err != nil {
				return ctx.SendString("遇到错误")
			}
			return nil
		},
	})
	//全局的错误处理
	app.Use(recover.New())

	app.Get("/demo", func(ctx *fiber.Ctx) error {
		panic("失败了")
	})

	app.Get("/getUser", func(ctx *fiber.Ctx) error {
		validate := validator.New()
		u := new(User)
		err2 := ctx.BodyParser(u)
		fmt.Println(err2)
		err := validate.Struct(u)
		fmt.Println(err)
		return ctx.SendString("校验")
	})

	log.Fatalln(app.Listen(":8002"))

}

type User struct {
	Name     string `validate:"required,min=3,max=32" json:"name"`
	IsActive *bool  `validate:"required" json:"isActive"`
}
