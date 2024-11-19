package handler

import (
	"net/http"
	"time"

	"github.com/Zen-Capital/nexa-backend/internal/model"
	"github.com/Zen-Capital/nexa-backend/internal/service"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	type UserInput struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var input UserInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Input Inválido",
		})
	}

	if len(input.Password) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Senha precisa ter no mínimo 6 caracteres",
		})
	}

	if input.Username == "" || input.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Todos os campos precisam ser prenchidos",
		})
	}

	// Encriptação da Senha //
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao criptografar a senha",
		})
	}

	// Criação do Objeto Usuário //
	user := model.User{
		ID:         primitive.NewObjectID(),
		Email:      input.Email,
		Username:   input.Username,
		Password:   string(hashedPassword),
		Account:    model.Account{},
		IsActive:   true,
		DateJoined: time.Now(),
	}

	if err := service.SaveUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao salvar usuário",
		})

	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Usuário criado com sucesso",
	})
}
