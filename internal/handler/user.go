package handler

import (
	"time"

	"github.com/Zen-Capital/nexa-backend/internal/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func createUser(c *fiber.Ctx) error {
	type UserInput struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var input UserInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"erro": "Input Inválido",
		})
	}

	if len(input.Password) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"erro": "Senha precisa ter no mínimo 6 caracteres",
		})
	}

	if input.Username == "" || input.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"erro": "Todos os campos precisam ser prenchidos",
		})
	}

	// Encriptação da Senha //
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"erro": "Erro ao encriptar senha",
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
			"erro": "Erro ao salvar usuário",
		})
	}

	// Geração de Token JWT //

}
