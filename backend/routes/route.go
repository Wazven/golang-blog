package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wazven/backendblog/controller"
	"github.com/wazven/backendblog/middleware"
)

func Setup(app *fiber.App){

	blog := app.Group("/blog")

	blog.Post("/blog/register", controller.Register)
	blog.Post("/blog/login", controller.Login)
	//middleware
	blog.Use(middleware.IsAuthenticate)
	blog.Post("/blog/post", controller.CreatePost)
	blog.Get("/blog/post", controller.GetAllPost)
	blog.Get("/blog/post/:id", controller.DetailPost)
	blog.Put("/blog/post/:id", controller.UpdatePost)
	blog.Get("/blog/uniquepost", controller.UniquePost)
	blog.Delete("/blog/post/:id", controller.DeletePost)
	blog.Post("blog/imagepost", controller.Upload)
	blog.Static("/blog/uploads", "./uploads")
}