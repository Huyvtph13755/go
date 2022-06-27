package controllers

import (
	"github.com/Huyvtph13755/go/config"
	"github.com/Huyvtph13755/go/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	type Params struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	params := new(Params)
	if err := c.BodyParser(params); err != nil {
		return err
	}
	session, err := config.SessionStore.Get(c)
	if err != nil {
		return err
	}
	if email := session.Get("email"); email != nil {
		return c.Status(422).JSON("Bạn đã đăng nhập")
	}
	var user *models.User
	result := config.Database.First(&user, "email = ?", params.Email)
	if result.Error != nil {
		return c.Status(401).JSON("Không tìm thấy user")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		return c.Status(401).JSON("Mật khẩu sai")
	}

	session.Set("email", user.Email)
	if err := session.Save(); err != nil {
		return c.Status(500).JSON("Lỗi lưu session")
	}
	return c.Status(200).JSON("Đăng nhập thành công")
}
func Register(c *fiber.Ctx) error {
	type Params struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	params := new(Params)
	if err := c.BodyParser(params); err != nil {
		return err
	}
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &models.User{
		Email:    params.Email,
		Password: string(hashed_password),
	}
	config.Database.Create(&user)
	return c.Status(201).JSON(user)
}

func Logout(c *fiber.Ctx) error {
	session, err := config.SessionStore.Get(c)
	if err != nil {
		return err
	}
	if err := session.Destroy(); err != nil {
		return c.Status(500).JSON("Lỗi xóa session")
	}
	return c.Status(200).JSON("Đăng xuất thành công")
}

// func GetUserByParam(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")
// 	if err != nil {
// 		return errors.New("invalid ID")
// 	}

// 	var user *models.User
// 	config.Database.Preload("Order").First(&user, id)

// 	return c.JSON(user)
// }

// func CreateUser(c *fiber.Ctx) error {
// 	type Params struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 		Role     string `json:"role"`
// 	}

// 	params := new(Params)
// 	if err := c.BodyParser(params); err != nil {
// 		return err
// 	}

// 	user := &models.User{
// 		Email:    params.Email,
// 		Password: params.Password,
// 		Role:     params.Role,
// 	}
// 	config.Database.Create(user)

// 	return c.Status(201).JSON(user)
// }

// func UpdateUser(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")
// 	if err != nil {
// 		return errors.New("invalid ID")
// 	}

// 	var user *models.User
// 	config.Database.First(&user, id)

// 	type Params struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 		Role     string `json:"role"`
// 	}

// 	params := new(Params)
// 	if err := c.BodyParser(params); err != nil {
// 		return err
// 	}

// 	config.Database.Model(&user).Updates(models.User{
// 		Email:    params.Email,
// 		Password: params.Password,
// 		Role:     params.Role,
// 	})

// 	return c.Status(201).JSON(user)
// }
// func DeleteUser(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")
// 	if err != nil {
// 		return errors.New("invalid ID")
// 	}

// 	var user *models.User
// 	config.Database.Where("id = ?", id).Delete(&user)
// 	return c.Status(201).JSON(user)
// }
