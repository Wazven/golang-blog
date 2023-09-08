package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wazven/backendblog/controller"
	"github.com/wazven/backendblog/middleware"
)

func Setup(app *fiber.App){
	app.Post("/blog/register", controller.Register)
	app.Post("/blog/login", controller.Login)
	//middleware
	app.Use(middleware.IsAuthenticate)
	app.Post("/blog/post", controller.CreatePost)
	app.Get("/blog/getallpost", controller.GetAllPost)
}