package controller

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/wazven/backendblog/database"
	"github.com/wazven/backendblog/models"
	"github.com/wazven/backendblog/util"
)

func ValidateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9. _%+\-]+@[a-z0-9. _%+\-]+\.[a-z0-9. _%+\-]`)
	return Re.MatchString(email)
}

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User
	if err := c.BodyParser(&data);err != nil{
		fmt.Println("Unable to parse body")
	}
	//Check Password apabila kurang dari 6 Karakter
	if len(data["password"].(string)) <=6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"Password harus lebih dari 6 karakter",
		})
	}

	if !ValidateEmail(strings.TrimSpace(data["email"].(string))){
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"Alamat Email Salah",
		})
	}

	//Check apabila email telah ada di database
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.ID != 0{
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"Email Sudah Ada, Silahkan Check Password Anda",
		})
	}
	user:=models.User{
		FirstName: data["first_name"].(string),
		LastName: data["last_name"].(string),
		Phone: data["phone"].(string),
		Email: strings.TrimSpace(data["email"].(string)),
	}

	user.SetPassword(data["password"].(string))
	err:=database.DB.Create(&user)
	if err != nil{
		log.Println(err)
	}
	c.Status(200)
		return c.JSON(fiber.Map{
			"user":user,
			"message":"Akun Berhasil Dibuat",
		})

}

func Login(c *fiber.Ctx)error {
	var data map[string]string
	if err:=c.BodyParser(&data);err!=nil{
		fmt.Println("Unable to parse body")
	}
	var user models.User
	database.DB.Where("email=?", data["email"]).First(&user)
	if user.ID ==0{
		c.Status(404)
		return c.JSON(fiber.Map{
			"message":"Alamat Email Tidak Ditemukan",
		})
	}
	if err:= user.ComparePassword(data["password"]); err !=nil{
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"Password Salah",
		})
	}
	token, err:= util.TokenJwt(strconv.Itoa(int(user.ID)),)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	cookie := fiber.Cookie{
		Name:"jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour *24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message":"Berhasil Login",
		"user":user,
		"token": token,
	})

}

type Claims struct{
	jwt.StandardClaims
}