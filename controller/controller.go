package controller

import (
	"fmt"
	"startupdigital/database"
	"startupdigital/model"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password := fmt.Sprintf("%v", data["password"])

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	name := fmt.Sprintf("%v", data["name"])
	email := fmt.Sprintf("%v", data["email"])
	phone := fmt.Sprintf("%v", data["phone"])
	var jk uint = uint(data["jk"].(float64))
	var role uint = 1
	var domisili uint = uint(data["domisili"].(float64))
	var kota_pelak uint = uint(data["kota_pelak"].(float64))

	user := model.User{
		Name:       name,
		Phone:      phone,
		Jk:         jk,
		Role:       role,
		Domisili:   domisili,
		Kota_pelak: kota_pelak,
		Email:      email,
		Password:   passwordHash,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user model.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Login Success",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user model.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func JawabTest(c *fiber.Ctx) error {
	data := map[string]interface{}{}
	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	var soal uint = uint(data["soal"].(float64))
	var jawaban uint = uint(data["jawaban"].(float64))

	JawabTest := model.JawabTest{
		Soal:    soal,
		Jawaban: jawaban,
	}

	database.DB.Create(&JawabTest)

	return c.JSON(fiber.Map{
		"message": "Success Menjawab",
	})
}
