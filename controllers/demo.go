package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"log"
)

func main() {
	//template
	engine := html.New("template", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Static("/static", "template/static")

	app.Get("/main", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{"username": "heqin"})
	})

	fun01 := func(ctx *fiber.Ctx) error {
		fmt.Println("经过中间件01")
		return ctx.Next()
	}

	fun02 := func(ctx *fiber.Ctx) error {
		fmt.Println("经过中间件02")
		return ctx.Next()
	}

	//增加一个中间件，第一个为路径匹配参数，但是不是完全匹配 例如 /demo01也会匹配到
	app.Use("/demo", fun01)
	app.Use("/api", fun02)

	//路由组
	group := app.Group("/test")
	group.Get("/a", func(ctx *fiber.Ctx) error {
		return ctx.SendString("test/a")
	})
	group.Get("/b", func(ctx *fiber.Ctx) error {
		return ctx.SendString("test/b")
	})

	//简单路由 如果同时符合通配符的规则，优先精确匹配的路由
	app.Get("/api/demo01", func(ctx *fiber.Ctx) error {

		return ctx.SendString("demo01")
	})

	//通配符
	app.Get("/api/*", func(ctx *fiber.Ctx) error {

		return ctx.SendString("通配符")

	})

	//贪婪匹配 +必须且任意匹配
	app.Get("/index/+", func(ctx *fiber.Ctx) error {
		return ctx.SendString("/index/+")
	})

	app.Get("/a/:name?", func(ctx *fiber.Ctx) error {
		return ctx.SendString("/api/:name")

	})

	log.Fatal(app.Listen(":8002"))
}
