package service

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/helper"
	"api-fiber-gorm/model"
	"api-fiber-gorm/request"
	"api-fiber-gorm/response"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthLogin(c *fiber.Ctx, request *request.Login) error {
	user, err := new(model.User), *new(error)

	if helper.ValidEmail(request.Identity) == true {
		user, err = findByEmail(request.Identity)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": false, "message": "Email not found !", "data": nil})
		}
	} else {
		user, err = findByUsername(request.Identity)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": false, "message": "Username not found !", "data": nil})
		}
	}

	if user == nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "User not found !", "data": nil})
	}

	if !helper.CheckPasswordHash(request.Password, user.Password) {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Invalid password !", "data": nil})
	}

	refreshToken, err := createRefreshToken(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	access_token, err := helper.CreateJwtToken(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	refresh_token, err := helper.CreateRefreshJwtToken(&refreshToken)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	var response = struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}{access_token, refresh_token}

	return c.JSON(fiber.Map{"status": true, "message": "success", "data": response})
}

func AuthMe(c *fiber.Ctx) error {
	tokenIdentity, err := helper.JwtParse(c)
	if len(err) > 1 {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	db := database.DB

	var user model.User
	if err := db.First(&user, "id = ?", tokenIdentity).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "User not found !", "data": nil})
	}

	var role model.Role
	if err := db.First(&role, "id = ?", user.RoleId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Role not found !", "data": nil})
	}

	permissions, err := helper.Permissions(user.RoleId)
	if len(err) > 1 {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	response := response.UserRolePermissionResponse(&user, &role, permissions)

	return c.JSON(fiber.Map{"status": true, "message": "success", "data": response})
}

func AuthRegister(c *fiber.Ctx, request *request.Register) error {
	if !helper.ValidEmail(request.Email) {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Email not valid !", "data": nil})
	}

	password, err := helper.HashPassword(request.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	user := model.User{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Password: password,
	}

	db := database.DB
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	user.Password = ""

	return c.JSON(fiber.Map{"status": true, "message": "success", "data": user})
}

func AuthRefreshToken(c *fiber.Ctx, request *request.RefreshToken) error {
	tokenIdentity, err := helper.RefreshTokenParse(c, request.RefreshToken)
	if len(err) > 1 {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	db := database.DB

	var refreshToken model.RefreshToken
	if err := db.First(&refreshToken, "id = ?", tokenIdentity).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "User not found !", "data": nil})
	}

	var user model.User
	if err := db.First(&user, "id = ?", refreshToken.UserId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "User not found !", "data": nil})
	}

	data := model.RefreshToken{
		Revoked:   false,
		ExpiredAt: time.Now().AddDate(0, 0, 2),
		// ExpiredAt: time.Now().Add(time.Minute * 120).Unix(),
	}

	if err := db.Model(&refreshToken).Updates(&data).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	access_token, _ := helper.CreateJwtToken(&user)
	refresh_token, _ := helper.CreateRefreshJwtToken(&refreshToken)

	var response = struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}{access_token, refresh_token}

	return c.JSON(fiber.Map{"status": true, "message": "success", "data": response})
}

func createRefreshToken(user *model.User) (model.RefreshToken, error) {
	refreshToken := model.RefreshToken{
		Revoked:   false,
		ExpiredAt: time.Now().AddDate(0, 0, 2),
		// ExpiredAt: time.Now().Add(time.Minute * 120).Unix(),
		UserId: user.Id,
	}

	db := database.DB
	if err := db.Create(&refreshToken).Error; err != nil {
		return refreshToken, err
	}

	return refreshToken, nil
}

func findByEmail(email string) (*model.User, error) {
	db := database.DB
	var user model.User
	if err := db.Where(&model.User{Email: email}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func findByUsername(username string) (*model.User, error) {
	db := database.DB
	var user model.User
	if err := db.Where(&model.User{Username: username}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
