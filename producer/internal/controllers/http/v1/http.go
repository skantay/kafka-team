package v1

import (
	"github.com/gofiber/fiber/v3"
	"github.com/skantay/service/internal/entities"
	"github.com/skantay/service/internal/usecase"
)

type Controller struct {
	usecase usecase.Usecase
}

func Run(usecase usecase.Usecase) {
	app := fiber.New()

	app.Post("/create", func(c fiber.Ctx) error {
		return usecase.Produce(entities.Entity{
			Name:      "hello",
			Age:       15,
			Relatives: []string{"1", "2", "3"},
		})
	})

	app.Listen(":8081")
}
