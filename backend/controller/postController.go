package controller

import (
	"fmt"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/wazven/backendblog/database"
	"github.com/wazven/backendblog/models"
)

func CreatePost(c *fiber.Ctx) error {
	var postBlog models.Blog
	if err := c.BodyParser(&postBlog);err != nil{
		fmt.Println("Unable to parse body")
	}

	if err := database.DB.Create(&postBlog).Error; err!=nil{
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"Invalid Payload",
		})
	}
	return c.JSON(fiber.Map{
		"message":"Berhasil Posting",
	})
}

func GetAllPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("Halaman", "1"))
	limit := 5
	offset := (page-1) * limit
	var total int64
	var getallblog []models.Blog
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getallblog)
	database.DB.Model(&models.Blog{}).Count(&total)

	return c.JSON(fiber.Map{
		"data":getallblog,
		"meta":fiber.Map{
			"total":total,
			"page":page,
			"last_page":math.Ceil(float64(int(total)/limit)),
		},
	})
}

func DetailPost(c *fiber.Ctx) error {
	id, _ :=strconv.Atoi(c.Params("id"))
	var postBlog models.Blog
	database.DB.Where("id=?", id).Preload("User").First(&postBlog)
	return c.JSON(fiber.Map{
		"data":postBlog,
	})
}

func UpdatPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		ID: uint(id),
	}
	if err := c.BodyParser(&blog);err!=nil {
		fmt.Println("Unable to parse body")
	}
	database.DB.Model(&blog).Updates(blog)
	return c.JSON(fiber.Map{
		"message":"Edit Posting Berhasil",
	})
}
